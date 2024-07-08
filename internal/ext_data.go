package internal

import (
	"github.com/HEUDavid/go-fsm-demo/model"
)

type MyExtData struct {
	model.ExtData
}

func (d MyExtData) SetTaskID(taskID string) {
	d.TaskID = taskID
}
