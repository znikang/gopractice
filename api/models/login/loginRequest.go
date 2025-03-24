package login

import (
	jwt5 "github.com/golang-jwt/jwt/v5"
)

var LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var JwtSecret string
var RefshToeknSecret string

type Claims struct {
	Username string `json:"username"`
	jwt5.RegisteredClaims
}
