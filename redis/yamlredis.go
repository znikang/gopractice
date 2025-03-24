package yamlredis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"log"
	"strconv"
	"yaml/common"
)

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	DB       string `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

func InitializeRedis() *redis.Client {
	szaddr := fmt.Sprintf(" %s:%d", common.Bargconfig.Redis.Host, common.Bargconfig.Redis.Port)
	log.Printf(szaddr + " (" + common.Bargconfig.Redis.Password + ")" + fmt.Sprintf("%d", common.Bargconfig.Redis.DB))
	client := redis.NewClient(&redis.Options{
		//		Addr:     szaddr,
		Addr:     common.Bargconfig.Redis.Host + ":" + strconv.Itoa(common.Bargconfig.Redis.Port),
		Password: common.Bargconfig.Redis.Password, // no password set
		DB:       common.Bargconfig.Redis.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}
