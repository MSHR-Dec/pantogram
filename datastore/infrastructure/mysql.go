package infrastructure

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMysql() (db *gorm.DB) {
	endpoint := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		environment.MysqlUser,
		environment.MysqlPassword,
		environment.MysqlUrl,
		environment.MysqlDbname)

	db, err := gorm.Open(mysql.Open(endpoint), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if environment.Env != "prd" {
		db.Logger = logger.Default.LogMode(logger.Info)
	}

	return
}
