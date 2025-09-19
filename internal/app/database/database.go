package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
var DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func init() {
	if err != nil {
		panic("数据库连接失败")
	}
}
