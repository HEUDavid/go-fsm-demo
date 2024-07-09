package internal

import (
	"context"
	"github.com/HEUDavid/go-fsm-demo/model"
	"github.com/HEUDavid/go-fsm/pkg"
	db "github.com/HEUDavid/go-fsm/pkg/db/mysql"
	. "github.com/HEUDavid/go-fsm/pkg/metadata"
	mq "github.com/HEUDavid/go-fsm/pkg/mq/rmq"
	"github.com/HEUDavid/go-fsm/pkg/util"
	"log"
	"sync"
	"time"
)

type MyAdapter struct {
	pkg.Adapter[*MyExtData]
}

func (a *MyAdapter) BeforeCreate(c context.Context, task *Task[*MyExtData]) error {
	log.Println("Rewrite BeforeCreate...")
	task.Version = 1
	task.ExtData.TransactionTime = time.Now().UnixNano() / int64(time.Millisecond)
	return nil
}

func NewMyAdapter() *MyAdapter {
	a := &MyAdapter{}
	a.ReBeforeCreate = a.BeforeCreate
	return a
}

var Adapter = NewMyAdapter()
var adapterInit sync.Once

func init() {
	adapterInit.Do(func() {
		Adapter.RegisterModel(
			&MyExtData{},
			&model.Task{},
			&model.UniqueRequest{},
		)
		Adapter.RegisterFSM(PayFsm)
		Adapter.RegisterDB(&db.Factory{})
		Adapter.RegisterMQ(&mq.Factory{})
		Adapter.Config = util.GetConfig()
		_ = Adapter.Init()
	})
}
