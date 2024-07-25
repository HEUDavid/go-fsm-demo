package internal

import (
	"github.com/HEUDavid/go-fsm-demo/model"
)

type MyData struct {
	model.Data
}

func (m *MyData) TableName() string {
	return m.Data.TableName()
}

func (m *MyData) SetTaskID(taskID string) {
	m.TaskID = taskID
}
