package v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"webserver/common"
)

type TestData struct {
	Hello string `json:"hello"`
}

func Test2(c *gin.Context) {
	data := new(TestData)
	data.Hello = fmt.Sprintf("world 2 %+v", common.Bargconfig.Server.Host) // %+v")
	c.JSON(http.StatusOK, data)
}
