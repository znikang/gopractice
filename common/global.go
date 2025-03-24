package common

import (
	"sync"
	"time"
	"yaml/common/config"
)

const (
	// Version go-admin version info
	Version = "1.0.0"
)

var (
	Bargconfig config.BuConfig
	once       sync.Once
)

const AccessTokenExpire = time.Minute * 15
const RefreshTokenExpire = time.Hour * 24 * 7

func init() {
	Bargconfig.Server.Port = 1234
	Bargconfig.Server.Host = "127.0.0.1"

	Bargconfig.Redis.Port = 6379
	Bargconfig.Redis.Host = "127.0.0.1"
	Bargconfig.Redis.Password = ""
	Bargconfig.Redis.DB = 0
}
