package internal

import (
	"github.com/HEUDavid/go-fsm-demo/model"
	"github.com/HEUDavid/go-fsm/pkg"
	db "github.com/HEUDavid/go-fsm/pkg/db/mysql"
	mq "github.com/HEUDavid/go-fsm/pkg/mq/rmq"
	"github.com/HEUDavid/go-fsm/pkg/util"
	"sync"
)

type ServiceWorker struct {
	pkg.Worker[*MyData]
	__init__ sync.Once
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
		s.RegisterDB(&db.Factory{Section: "mysql_public"})
		s.RegisterMQ(&mq.Factory{Section: "rmq_public"})
		s.Config = util.GetConfig()
		s.Init()
	})
}

func NewWorker() *ServiceWorker {
	w := &ServiceWorker{}
	w.MaxGoroutines = 50
	return w
}
