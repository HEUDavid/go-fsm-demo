package internal

import (
	"github.com/HEUDavid/go-fsm-demo/model"
	. "github.com/HEUDavid/go-fsm/pkg/metadata"
)

type MyExtData struct {
	model.ExtData
	ExtDataEntity
}

func (d MyExtData) TableName() string {
	return "ext_data"
}

func (d MyExtData) SetTaskID(taskID string) {
	d.TaskID = taskID
}
