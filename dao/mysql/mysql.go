package mysql

import (
	"bookSystem/model"
	"bookSystem/settings"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitMysqlDB(cfg *settings.MysqlConf) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Db)
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(errors.New("连接数据库失败, " + err.Error()))
		return errors.New("连接数据库失败, " + err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = db
	db.AutoMigrate(&model.User{}, &model.Book{})

	return nil

}
