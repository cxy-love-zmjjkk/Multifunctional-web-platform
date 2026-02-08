package config

import (
	"exchange_backend/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func initDB() {
	dsn := AppConfig.Database.Dsn

	//与数据库建立连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize database,got err:%v", err)
	}
	//配置
	sqlDB, err := db.DB()

	//设置空闲连接的最大数量
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)

	//设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)

	//每个连接可复用的最大时间,过了这个时间就得重新建立连接
	sqlDB.SetConnMaxLifetime(time.Hour)

	//错误处理
	if err != nil {
		log.Fatalf("Failed to configure database,got error:%v", err)
	}
	global.Db = db
}
