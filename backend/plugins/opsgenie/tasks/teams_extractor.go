/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tasks

import (
	"encoding/json"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/opsgenie/models"
	"github.com/apache/incubator-devlake/plugins/opsgenie/models/raw"
)

var _ plugin.SubTaskEntryPoint = ExtractTeams

var ExtractTeamsMeta = plugin.SubTaskMeta{
	Name:             "extractTeams",
	EntryPoint:       ExtractTeams,
	EnabledByDefault: true,
	Description:      "extract Opsgenie teams",
	DomainTypes:      []string{plugin.DOMAIN_TYPE_CROSS},
}

func ExtractTeams(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*OpsgenieTaskData)
	extractor, err := api.NewApiExtractor(api.ApiExtractorArgs{
		RawDataSubTaskArgs: api.RawDataSubTaskArgs{
			Ctx:     taskCtx,
			Options: data.Options,
			Table:   RAW_TEAMS_TABLE,
		},
		Extract: func(row *api.RawData) ([]interface{}, errors.Error) {
			teamRaw := &raw.Team{}
			err := errors.Convert(json.Unmarshal(row.Data, teamRaw))
			if err != nil {
				return nil, err
			}

			results := make([]interface{}, 0, 1)
			user := models.Team{
				ConnectionId: data.Options.ConnectionId,
				Id:           *teamRaw.Id,
				Name:         *teamRaw.Name,
				Description:  *teamRaw.Description,
			}
			results = append(results, &user)
			return results, nil
		},
	})
	if err != nil {
		return err
	}
	return extractor.Execute()
}