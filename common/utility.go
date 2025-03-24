package common

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	RedisCli *redis.Client
	OrmCli   *gorm.DB
)
