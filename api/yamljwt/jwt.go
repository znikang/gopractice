package yamljwt

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	jwt4 "github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
	"webserver/common/vo"
)

var (
	identityKey = "id"
	port        string
)

func InitParams() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: payloadFunc(),

		IdentityHandler: identityHandler(),
		Authenticator:   authenticator(),
		Authorizator:    authorizator(),
		Unauthorized:    unauthorized(),
		LoginResponse:   LoginResponse(),
		LogoutResponse:  LogoutResponse(),
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}

func HandlerMiddleWare(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	fmt.Println("HandlerMiddleWare")
	return func(context *gin.Context) {
		fmt.Println("HandlerMiddleWare 3")
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}

func identityHandler() func(c *gin.Context) interface{} {
	fmt.Println("identityHandler")
	return func(c *gin.Context) interface{} {
		fmt.Printf("identityHandler 1 %+v ", c)

		claims := jwt.ExtractClaims(c)
		fmt.Printf("identityHandler 1 %+v ", claims)
		return &vo.User{
			UserName: claims[identityKey].(string),
		}
	}
}

func payloadFunc() func(data interface{}) jwt4.MapClaims {
	fmt.Printf("payloadFunc 1 ")
	return func(data interface{}) jwt4.MapClaims {
		fmt.Println("payloadFunc 4 ")
		if v, ok := data.(*vo.User); ok {
			return jwt4.MapClaims{
				identityKey: v.UserName,
			}
		}
		return jwt4.MapClaims{}
	}
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	fmt.Println("authenticator")
	return func(c *gin.Context) (interface{}, error) {
		fmt.Println("authenticato 2r")
		var loginVals vo.Login
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		userID := loginVals.UserName
		password := loginVals.PassWord

		if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
			return &vo.User{
				UserName:  userID,
				LastName:  "Bo-Yi",
				FirstName: "Wu",
			}, nil
		}
		return nil, jwt.ErrFailedAuthentication
	}
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	fmt.Println("authorizator")
	return func(data interface{}, c *gin.Context) bool {
		fmt.Println("authorizator 3")
		if v, ok := data.(*vo.User); ok && v.UserName == "admin" {
			return true
		}
		return false
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	fmt.Println("unauthorized")
	return func(c *gin.Context, code int, message string) {
		fmt.Println("unauthorized 4")
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}

func LoginResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	fmt.Println("LoginResponse")
	return func(c *gin.Context, code int, token string, expire time.Time) {
		fmt.Println("LoginResponse 2")
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		})
	}
}

func LogoutResponse() func(c *gin.Context, code int) {
	return func(c *gin.Context, code int) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})
	}
}

func HandleNoRoute() func(c *gin.Context) {
	fmt.Println("handleNoRoute")
	return func(c *gin.Context) {
		fmt.Println("handleNoRoute 5")
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	}
}

func HelloHandler(c *gin.Context) {
	fmt.Println("helloHandler")
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*vo.User).UserName,
		"text":     "Hello World.",
	})
}
