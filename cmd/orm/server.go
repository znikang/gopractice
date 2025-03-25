package ormserver

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
	"time"
	"webserver/database/orm"
)

var (
	StartCmd = &cobra.Command{
		Use:     "gorm",
		Short:   "gorm version info",
		Example: "gorm version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

type Usermysql struct {
	//gorm為model的tag標籤，v2版的auto_increment要放在type裡面，v1版是放獨立定義
	ID        int64     `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Username  string    `gorm:"type:varchar(20) NOT NULL;" json:"username,omitempty"`
	Password  string    `gorm:"type:varchar(100) NOT NULL;" json:"password,omitempty"`
	Status    int32     `gorm:"type:int(5);" json:"status,omitempty"`
	CreatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

var dbselect string

func init() {
	StartCmd.Flags().StringVar(&dbselect, "database", "", "Start server with provided dbtabase")
}

const (
	UserName     string = "root"
	Password     string = "1qaz@WSX"
	Addr         string = "192.168.1.215"
	Port         int    = 3306
	Database     string = "g_paypay"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func mysqloutput() {

	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", UserName, Password, Addr, Port, Database)
	db, _ := gorm.Open(mysql.Open(addr), &gorm.Config{})

	// 初始化 GORM 生成器
	g := gen.NewGenerator(gen.Config{
		OutPath: "./model", // 輸出目錄
	})

	// 設定 GORM 數據庫
	g.UseDB(db)
	// 自動生成 Model（根據資料表）
	g.GenerateAllTable()

	// 寫入檔案
	g.Execute()
}
func mysqlf() {
	fmt.Println("mysqlf")
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", UserName, Password, Addr, Port, Database)
	conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	db, err1 := conn.DB()
	if err1 != nil {
		fmt.Println("get db failed:", err)
		return
	}
	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)
	err = conn.AutoMigrate(&Usermysql{})
	if err != nil {
		log.Fatal("❌ AutoMigrate 失敗:", err)
	}
	err = conn.AutoMigrate(&orm.Acl{})
	if err != nil {
		log.Fatal("❌ AutoMigrate 失敗:", err)
	}
	//migrator := conn.Migrator()
	//	has := migrator.HasTable(&orm.User{})
	//has := migrator.HasTable("GG")
	//	if !has {
	//		fmt.Println("table not exist")
	//	}
}
func sqlitef() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&orm.User{})
	db.Create(&orm.User{Name: "Alice", Age: 25})
	db.Create(&orm.User{Name: "Bob", Age: 30})

	var user orm.User
	db.First(&user, 1) // 透過 ID 查找
	fmt.Println("User 1:", user)

	// 查詢所有使用者
	var users []orm.User
	db.Find(&users)
	fmt.Println("All Users:", users)

	// 更新使用者
	db.Model(&user).Update("Age", 26)
}
func run() error {

	fmt.Println("gorm test")
	if dbselect == "mysql" {
		fmt.Println("mysql")
		mysqlf()
	} else if dbselect == "mysqlout" {
		fmt.Println("mysqlout")
		mysqloutput()
	} else if dbselect == "sqlite" {
		fmt.Println("sqlite")
		sqlitef()
	}

	return nil
}
