package main

import (
	"context"
	"github.com/HEUDavid/go-fsm-demo/internal"
	. "github.com/HEUDavid/go-fsm-demo/internal/pkg"
	"github.com/HEUDavid/go-fsm-demo/model"
	"github.com/HEUDavid/go-fsm/pkg"
	db "github.com/HEUDavid/go-fsm/pkg/db/mysql"
	. "github.com/HEUDavid/go-fsm/pkg/metadata"
	mq "github.com/HEUDavid/go-fsm/pkg/mq/rmq"
	"github.com/HEUDavid/go-fsm/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"sync"
	"time"
)

type MyAdapter struct {
	pkg.Adapter[*internal.MyExtData]
}

func (a *MyAdapter) BeforeCreate(c context.Context, task *Task[*internal.MyExtData]) error {
	log.Println("Rewrite BeforeCreate...")
	task.Version = 1
	task.ExtData.TransactionTime = time.Now().UnixNano() / int64(time.Millisecond)
	return nil
}

func NewMyAdapter() *MyAdapter {
	_a := &MyAdapter{}
	_a.ReBeforeCreate = _a.BeforeCreate
	return _a
}

var adapter = NewMyAdapter()
var adapterInit sync.Once

type MyWorker struct {
	pkg.Worker[*internal.MyExtData]
}

func NewMyWorker() *MyWorker {
	_w := &MyWorker{}
	return _w
}

var worker = NewMyWorker()
var workerInit sync.Once

func init() {
	adapterInit.Do(func() {
		adapter.RegisterModel(
			&internal.MyExtData{},
			&model.Task{},
			&model.UniqueRequest{},
		)
		adapter.RegisterFSM(internal.PayFsm)
		adapter.RegisterDB(&db.Factory{})
		adapter.RegisterMQ(&mq.Factory{})
		adapter.Config = util.GetConfig()
		_ = adapter.Init()

	})
	workerInit.Do(func() {
		worker.RegisterModel(
			&internal.MyExtData{},
			&model.Task{},
			&model.UniqueRequest{},
		)
		worker.RegisterFSM(internal.PayFsm)
		worker.RegisterDB(&db.Factory{})
		worker.RegisterMQ(&mq.Factory{})
		worker.Config = util.GetConfig()
		worker.Init()
	})
}

func Create(c *gin.Context) {
	task := NewTaskInstance(
		c.Query("request_id"), "",
		&internal.MyExtData{ExtData: model.ExtData{
			Symbol: "BTC", Quantity: 1, Amount: 64000, Operator: "user1", Comment: c.Query("comment"),
		}},
	)
	task.Type = c.Query("type")

	err := adapter.Create(c, task)
	Response(c, err, task)
}

func Query(c *gin.Context) {
	task := NewTaskInstance(c.Query("request_id"), c.Query("task_id"), &internal.MyExtData{})
	err := adapter.Query(c, task)
	Response(c, err, task)
}

func Update(c *gin.Context) {
	task := NewTaskInstance(
		c.Query("request_id"), c.Query("task_id"),
		&internal.MyExtData{ExtData: model.ExtData{
			Symbol: "ETH", Quantity: 2, Amount: 70000, Operator: "", Comment: c.Query("comment"),
		}},
	)
	task.Type = c.Query("type")
	task.State = "End"
	task.Version, _ = strconv.Atoi(c.Query("version"))

	task.SetSelectColumns([]string{"Quantity", "Operator"})
	task.SetOmitColumns([]string{"Amount", "Symbol"})

	err := adapter.Update(c, task)
	Response(c, err, task)
}

func main() {
	worker.Run()
	log.Println("worker started...")

	r := gin.Default()
	r.GET("/create", Create)
	r.GET("/query", Query)
	r.GET("/update", Update)
	_ = r.Run("127.0.0.1:8080")
}
