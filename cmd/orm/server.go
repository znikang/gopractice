package ormserver

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
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

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100"`
	Age  int
}

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
	StartCmd.PersistentFlags().StringVar(&dbselect, "database", "", "Start server with provided dbtabase")
}

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

func mysqlf() {

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
	conn.AutoMigrate(&Usermysql{})

	migrator := conn.Migrator()
	has := migrator.HasTable(&User{})
	//has := migrator.HasTable("GG")
	if !has {
		fmt.Println("table not exist")
	}
}
func sqlitef() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "Alice", Age: 25})
	db.Create(&User{Name: "Bob", Age: 30})

	var user User
	db.First(&user, 1) // 透過 ID 查找
	fmt.Println("User 1:", user)

	// 查詢所有使用者
	var users []User
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
	} else if dbselect == "sqlite" {
		fmt.Println("sqlite")
		sqlitef()
	}

	return nil
}
