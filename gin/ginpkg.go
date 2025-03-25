package ginpkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	myapi "webserver/api"
	v22 "webserver/api/v2"
	"webserver/cmd/authserver"
)

func protectedHandler(c *gin.Context) {
	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{"message": "Welcome!", "user": username})
}

func InitGin(serverport string) *gin.Engine {
	router := gin.Default()

	router.POST("login", authserver.LoginHandler)
	router.POST("logout", authserver.LogoutHandler)
	router.POST("refreshtoken", authserver.RefreshTokenHandler)

	// 當沒有路由的配置時 會走這個
	router.NoRoute(authserver.AuthMiddleware(), authserver.HandleNoRoute())
	// 這個需要注意順序router.Use 如果在 login 前面 那login 會被要求  token
	router.Use(authserver.AuthMiddleware())

	router.GET("test", myapi.PrintMessage)
	router.GET("protected", authserver.AuthMiddleware(), protectedHandler)
	router.GET("hello2", v22.Test2)

	auth := router.Group("/auth")
	auth.GET("hello2", v22.Test2)
	router.Run(serverport)

	//srv := &http.Server{
	//	Addr:    fmt.Sprintf("%s:%d", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
	//	Handler: sdk.Runtime.GetEngine(),
	//} 不一樣的啟動方式

	return router
}
