package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"webserver/common"
)

const (
	maxLifetime  int = 10
	maxOpenConns int = 10
	maxIdleConns int = 10
)

func InitializeDatabases() *gorm.DB {

	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", common.Bargconfig.Database.Username,
		common.Bargconfig.Database.Password,
		common.Bargconfig.Database.Host,
		common.Bargconfig.Database.Port,
		common.Bargconfig.Database.DB)
	conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{})

	dbsetting, err1 := conn.DB()
	if err1 != nil {
		fmt.Println("get db failed:", err)
		return nil
	}
	dbsetting.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	dbsetting.SetMaxIdleConns(maxOpenConns)
	dbsetting.SetMaxOpenConns(maxIdleConns)
	return conn
}
