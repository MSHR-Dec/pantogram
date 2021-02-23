package persistence

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMysql(env string, user string, password string, host string, name string) (db *gorm.DB) {
	endpoint := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		name)

	db, err := gorm.Open(mysql.Open(endpoint), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if env != "prd" {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	return
}
