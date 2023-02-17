package config

import (
	"go-micro.dev/v4/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var _db *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/oceanmicro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Info("数据库连接失败...")
		os.Exit(1)
	}
	_db = db
}

func DB() *gorm.DB {
	return _db
}
