package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yaml/common"
)

type TestData struct {
	Hello string `json:"hello"`
}

func Test1(c *gin.Context) {
	data := new(TestData)
	data.Hello = fmt.Sprintf("world %+v", common.Bargconfig.Server.Host) // %+v")
	c.JSON(http.StatusOK, data)
}

func Login(c *gin.Context) {
	data := new(TestData)
	data.Hello = fmt.Sprintf("world %+v", common.Bargconfig.Server.Host) // %+v")
	c.JSON(http.StatusOK, data)
}
