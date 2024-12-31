package main

import (
	"fmt"
	. "github.com/HEUDavid/go-fsm-demo/internal"
	"github.com/HEUDavid/go-fsm-demo/model"
	. "github.com/HEUDavid/go-fsm/pkg/metadata"
	"github.com/HEUDavid/go-fsm/pkg/util"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func Create(c *gin.Context) {
	task := GenTaskInstance(
		c.Query("request_id"), "",
		&MyData{Data: model.Data{
			Symbol: "BTC", Quantity: 1, Amount: 64000, Operator: "user", Comment: c.Query("comment"),
		}})
	task.Data.UID = 10001
	task.Type = c.Query("type")
	task.State = New.GetName()
	_response(c, Adapter.Create(c, task), task)
}

func Update(c *gin.Context) {
	task := GenTaskInstance(
		c.Query("request_id"), c.Query("task_id"),
		&MyData{Data: model.Data{
			Symbol: "ETH", Quantity: -2, Amount: -6000, Operator: "", Comment: c.Query("comment"),
		}})
	task.Type = c.Query("type")
	task.State = "End"
	version, _ := strconv.ParseUint(c.Query("version"), 10, 64)
	task.Version = uint(version)
	task.SetSelectColumns([]string{"Operator"})
	task.SetOmitColumns([]string{"Symbol", "Quantity"})

	_response(c, Adapter.Update(c, task), task)
}

func Query(c *gin.Context) {
	task := GenTaskInstance(c.Query("request_id"), c.Query("task_id"), &MyData{})
	_response(c, Adapter.Query(c, task), task)
}

func _response(c *gin.Context, err error, task interface{}) {
	if err == nil {
		c.JSON(http.StatusOK, task)
	} else {
		c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}
}

var (
	Worker  = NewWorker()
	Adapter = NewAdapter()
)

func setupLog() {
	logPath := filepath.Join(util.FindProjectRoot(), "log/water.log")

	if err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm); err != nil {
		panic(fmt.Sprintf("Failed to create log directory: %v", err))
	}

	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to open or create log file: %v", err))
	}

	mw := io.MultiWriter(os.Stdout, f)
	gin.DefaultWriter = mw
	log.SetOutput(mw)

	log.SetPrefix("[FSM-DEMO] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func init() {
	setupLog()

	Worker.DoInit()
	Adapter.DoInit()
}

func main() {
	Worker.Run()
	log.Println("[FSM] worker started...")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/create", Create)
	r.GET("/update", Update)
	r.GET("/query", Query)
	_ = r.Run("127.0.0.1:8080")
}
