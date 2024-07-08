package pkg

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pretty(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

func Response(c *gin.Context, err error, task interface{}) {
	if err == nil {
		c.JSON(http.StatusOK, &task)
	} else {
		c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
	}
}
