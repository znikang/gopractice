package Database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"yaml/common"
)

const (
	UserName     string = "root"
	Password     string = "1qaz@WSX"
	Addr         string = "192.168.1.171"
	Port         int    = 3306
	Database     string = "gormtest"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
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
	dbsetting.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	dbsetting.SetMaxIdleConns(MaxIdleConns)
	dbsetting.SetMaxOpenConns(MaxOpenConns)
	return conn
}
