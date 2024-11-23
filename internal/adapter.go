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

type ServiceAdapter struct {
	pkg.Adapter[*MyData]
	__init__ sync.Once
}

func (a *ServiceAdapter) BeforeCreate(c context.Context, task *Task[*MyData]) error {
	log.Println("[FSM] Rewrite BeforeCreate...")
	task.Version = 1
	task.Data.TransactionTime = uint64(time.Now().Unix())
	return nil
}

func (a *ServiceAdapter) DoInit() {
	a.__init__.Do(func() {
		a.RegisterModel(
			&MyData{},
			&model.Task{},
			&model.UniqueRequest{},
		)
		a.RegisterFSM(PayFSM)
		a.RegisterGenerator(util.UniqueID)
		a.RegisterDB(&db.Factory{Section: "mysql_public"})
		a.RegisterMQ(&mq.Factory{Section: "rmq_public"})
		a.Config = util.GetConfig()
		_ = a.Init()
	})
}

func NewAdapter() *ServiceAdapter {
	a := &ServiceAdapter{}
	a.ReBeforeCreate = a.BeforeCreate
	return a
}
