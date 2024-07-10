package internal

import (
	"github.com/HEUDavid/go-fsm-demo/model"
)

type MyData struct {
	model.Data
}

func (d *MyData) SetTaskID(taskID string) {
	d.TaskID = taskID
}
