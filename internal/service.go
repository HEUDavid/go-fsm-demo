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

type IService interface{ DoInit() }
type Service struct{ __init__ sync.Once }

func (s *Service) DoInit() { panic("not implemented") }

type ServiceWorker struct {
	Service
	pkg.Worker[*MyData]
}

func (s *ServiceWorker) DoInit() {
	s.__init__.Do(func() {
		s.RegisterModel(
			&MyData{},
			&model.Task{},
			&model.UniqueRequest{},
		)
		s.RegisterFSM(PayFSM)
		s.RegisterGenerator(util.UniqueID)
		s.RegisterDB(&db.Factory{Section: "mysql"})
		s.RegisterMQ(&mq.Factory{Section: "rmq"})
		s.Config = util.GetConfig()
		s.Init()
	})
}

func NewWorker() *ServiceWorker {
	w := &ServiceWorker{}
	w.MaxGoroutines = 50
	w.DEBUG = true
	return w
}

type ServiceAdapter struct {
	Service
	pkg.Adapter[*MyData]
}

func (s *ServiceAdapter) DoInit() {
	s.__init__.Do(func() {
		s.RegisterModel(
			&MyData{},
			&model.Task{},
			&model.UniqueRequest{},
		)
		s.RegisterFSM(PayFSM)
		s.RegisterGenerator(util.UniqueID)
		s.RegisterDB(&db.Factory{Section: "mysql"})
		s.RegisterMQ(&mq.Factory{Section: "rmq"})
		s.Config = util.GetConfig()
		_ = s.Init()
	})
}

func (s *ServiceAdapter) BeforeCreate(c context.Context, task *Task[*MyData]) error {
	log.Println("[FSM-DEMO] Rewrite BeforeCreate...")
	task.Version = 1
	task.Data.TransactionTime = uint64(time.Now().Unix())
	return nil
}

func NewAdapter() *ServiceAdapter {
	a := &ServiceAdapter{}
	a.ReBeforeCreate = a.BeforeCreate
	return a
}
