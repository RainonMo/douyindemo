package common

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/driver/mysql"
)

var Db *gorm.DB

func Mysql() {
	
	dsn := "root:1234@tcp(127.0.0.1:3306)/douyin_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Printf("mysql error: %s", err)
		return
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Printf("mysql error: %s", err)
	}

	// 设置空闲连接池中连接的最大数量
	sqlDb.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量
	sqlDb.SetMaxOpenConns(100)
	// 设置了连接可复用的最大时间
	sqlDb.SetConnMaxLifetime(time.Hour)

	Db = db
}