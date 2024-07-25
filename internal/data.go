package internal

import (
	"github.com/HEUDavid/go-fsm-demo/model"
)

type MyData struct {
	model.Data
}

func (m *MyData) SetTaskID(taskID string) {
	m.TaskID = taskID
}
