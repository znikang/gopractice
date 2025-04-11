package common

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"webserver/grpccli"
)

var (
	RedisCli *redis.Client
	OrmCli   *gorm.DB
	GrpcCli  *grpccli.Clients
)
