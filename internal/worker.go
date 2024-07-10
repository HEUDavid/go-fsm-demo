package internal

import (
	"github.com/HEUDavid/go-fsm-demo/model"
	"github.com/HEUDavid/go-fsm/pkg"
	db "github.com/HEUDavid/go-fsm/pkg/db/mysql"
	mq "github.com/HEUDavid/go-fsm/pkg/mq/rmq"
	"github.com/HEUDavid/go-fsm/pkg/util"
	"sync"
)

type MyWorker struct {
	pkg.Worker[*MyData]
}

func NewMyWorker() *MyWorker {
	w := &MyWorker{}
	return w
}

var Worker = NewMyWorker()
var _workerInit sync.Once

func WorkerInit() {
	_workerInit.Do(func() {
		Worker.RegisterModel(
			&MyData{},
			&model.Task{},
			&model.UniqueRequest{},
		)
		Worker.RegisterFSM(PayFSM)
		Worker.RegisterDB(&db.Factory{})
		Worker.RegisterMQ(&mq.Factory{})
		Worker.Config = util.GetConfig()
		Worker.Init()
	})
}
