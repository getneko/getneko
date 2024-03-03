package db

import (
	"getneko/structtypes"
	"time"

	"github.com/glebarez/sqlite"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	ORMDB *gorm.DB
	err   error
)

func init() {
	//连接数据库
	ORMDB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := ORMDB.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	//迁移数据库
	err := ORMDB.AutoMigrate(&structtypes.User{})
	if err != nil {
		log.Error("Database connection initialization failed", err.Error())
	}
	err = ORMDB.AutoMigrate(&structtypes.Project{})
	if err != nil {
		log.Error("Database connection initialization failed", err.Error())
	} else {
		log.Info("Database connection initialization successful")
	}
}
