package common

import (
	"sync"
	"time"
	"webserver/common/config"
)

const (
	// Version go-admin version info
	Version = "1.0.0"
)

var (
	Bargconfig config.BuConfig
	once       sync.Once
)

const AccessTokenExpire = time.Minute * 20
const RefreshTokenExpire = time.Hour

func init() {
	Bargconfig.Server.Port = 1234
	Bargconfig.Server.Host = "127.0.0.1"

	Bargconfig.Redis.Port = 6379
	Bargconfig.Redis.Host = "127.0.0.1"
	Bargconfig.Redis.Password = ""
	Bargconfig.Redis.DB = 0

	Bargconfig.Database.DB = "gormtest"
	Bargconfig.Database.Username = "admin"
	Bargconfig.Database.Username = "1qaz@WSX"
	Bargconfig.Database.Port = 3306
	Bargconfig.Database.Host = "192.168.1.171"
}
