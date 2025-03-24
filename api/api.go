package myapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestData struct {
	Hello string `json:"hello"`
}

func PrintMessage(c *gin.Context) {
	data := new(TestData)
	data.Hello = "world!"
	c.JSON(http.StatusOK, data)
}
