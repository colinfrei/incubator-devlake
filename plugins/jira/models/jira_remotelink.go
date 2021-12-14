package models

import (
	"github.com/merico-dev/lake/models/common"
	"gorm.io/datatypes"
)

type JiraRemotelink struct {
	common.NoPKModel

	// collected fields
	SourceId     uint64 `gorm:"primaryKey"`
	RemotelinkId uint64 `gorm:"primarykey"`
	IssueId      uint64 `gorm:"index"`
	RawJson      datatypes.JSON
	Self         string
	Title        string
	Url          string
}
