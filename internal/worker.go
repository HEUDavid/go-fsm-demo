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
	pkg.Worker[*MyExtData]
}

func NewMyWorker() *MyWorker {
	w := &MyWorker{}
	return w
}

var Worker = NewMyWorker()
var workerInit sync.Once

func init() {
	workerInit.Do(func() {
		Worker.RegisterModel(
			&MyExtData{},
			&model.Task{},
			&model.UniqueRequest{},
		)
		Worker.RegisterFSM(PayFsm)
		Worker.RegisterDB(&db.Factory{})
		Worker.RegisterMQ(&mq.Factory{})
		Worker.Config = util.GetConfig()
		Worker.Init()
	})
}
