package common

import (
	"sync"
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

func init() {
	Bargconfig.Server.Port = 1234
	Bargconfig.Server.Host = "127.0.0.1"

	Bargconfig.Redis.Port = 6379
	Bargconfig.Redis.Host = "127.0.0.1"
	Bargconfig.Redis.Password = ""
	Bargconfig.Redis.DB = 0
}
