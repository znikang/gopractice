package login

import (
	jwt5 "github.com/golang-jwt/jwt/v5"
	"yaml/common"
)

var LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var JwtSecret = []byte(common.Bargconfig.Server.Secrectkey)

type Claims struct {
	Username string `json:"username"`
	jwt5.RegisteredClaims
}
