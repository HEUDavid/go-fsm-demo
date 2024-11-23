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

func (w *ServiceWorker) DoInit() {
	w.__init__.Do(func() {
		w.RegisterModel(
			&MyData{},
			&model.Task{},
			&model.UniqueRequest{},
		)
		w.RegisterFSM(PayFSM)
		w.RegisterGenerator(util.UniqueID)
		w.RegisterDB(&db.Factory{Section: "mysql_public"})
		w.RegisterMQ(&mq.Factory{Section: "rmq_public"})
		w.Config = util.GetConfig()
		w.Init()
	})
}

func NewWorker() *ServiceWorker {
	w := &ServiceWorker{}
	w.MaxGoroutines = 50
	return w
}
