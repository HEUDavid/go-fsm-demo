package main

import (
	. "github.com/HEUDavid/go-fsm-demo/internal"
	"github.com/HEUDavid/go-fsm-demo/model"
	. "github.com/HEUDavid/go-fsm/pkg/metadata"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Create(c *gin.Context) {
	task := GenTaskInstance(
		c.Query("request_id"), "",
		&MyData{Data: model.Data{
			Symbol: "BTC", Quantity: 1, Amount: 64000, Operator: "user", Comment: c.Query("comment"),
		}})
	task.Type = c.Query("type")
	_response(c, Adapter.Create(c, task), task)
}

func Query(c *gin.Context) {
	task := GenTaskInstance(c.Query("request_id"), c.Query("task_id"), &MyData{})
	_response(c, Adapter.Query(c, task), task)
}

func Update(c *gin.Context) {
	task := GenTaskInstance(
		c.Query("request_id"), c.Query("task_id"),
		&MyData{Data: model.Data{
			Symbol: "ETH", Quantity: -2, Amount: -6000, Operator: "", Comment: c.Query("comment"),
		}})
	task.Type = c.Query("type")
	task.State = "End"
	task.Version, _ = strconv.Atoi(c.Query("version"))

	task.SetSelectColumns([]string{"Operator"})
	task.SetOmitColumns([]string{"Symbol", "Quantity"})

	_response(c, Adapter.Update(c, task), task)
}

func _response(c *gin.Context, err error, task interface{}) {
	if err == nil {
		c.JSON(http.StatusOK, &task)
	} else {
		c.JSON(http.StatusOK, map[string]error{"error": err})
	}
}

func main() {
	Worker.Run()
	log.Println("[FSM] Worker started...")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/create", Create)
	r.GET("/query", Query)
	r.GET("/update", Update)
	_ = r.Run("127.0.0.1:8080")
}
