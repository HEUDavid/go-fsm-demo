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
	pkg.Adapter[*MyData]
}

func (a *MyAdapter) BeforeCreate(c context.Context, task *Task[*MyData]) error {
	log.Println("[FSM] Rewrite BeforeCreate...")
	task.Version = 1
	task.Data.TransactionTime = uint64(time.Now().Unix())
	return nil
}

func NewMyAdapter() *MyAdapter {
	a := &MyAdapter{}
	a.ReBeforeCreate = a.BeforeCreate
	return a
}

var Adapter = NewMyAdapter()
var _initAdapter sync.Once

func InitAdapter() {
	_initAdapter.Do(func() {
		Adapter.RegisterModel(
			&MyData{},
			&model.Task{},
			&model.UniqueRequest{},
		)
		Adapter.RegisterFSM(PayFSM)
		Adapter.RegisterGenerator(util.UniqueID)
		Adapter.RegisterDB(&db.Factory{Section: "mysql_public"})
		Adapter.RegisterMQ(&mq.Factory{Section: "rmq_public"})
		Adapter.Config = util.GetConfig()
		_ = Adapter.Init()
	})
}
