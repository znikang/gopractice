package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/sharding"
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

func InitializeShardingDatabases() *gorm.DB {

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

	conn.Use(sharding.Register(sharding.Config{
		ShardingKey:         "user_id",
		NumberOfShards:      64,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "orders"))

	dbsetting.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	dbsetting.SetMaxIdleConns(maxOpenConns)
	dbsetting.SetMaxOpenConns(maxIdleConns)
	return conn
}

var shardDBs = map[int]*gorm.DB{}

func InitShardDBs() {
	for i := 0; i < 4; i++ {
		dsn := fmt.Sprintf("user:password@tcp(localhost:3306)/shard_%d?charset=utf8mb4&parseTime=True&loc=Local", i)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect to shard")
		}
		shardDBs[i] = db
	}
}

func GetShardDB(userID int) *gorm.DB {
	shardID := userID % 4
	return shardDBs[shardID]
}

//func CreateOrderByMonth(order *Order, year int, month int) error {
//	tableName := fmt.Sprintf("orders_%d_%02d", year, month)
//	return db.Table(tableName).Create(order).Error
//}
