package database

import (
	"app/envvars"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

func GormInit(DbEnv envvars.DbSetting) *GormAdapter {
	dbSetting := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&collation=%s&parseTime=%s",
		DbEnv.DB_USER,
		DbEnv.DB_PASSWORD,
		DbEnv.DB_PORT,
		DbEnv.DB_NAME,
		DbEnv.DB_CHARSET,
		DbEnv.DB_COLLATION,
		DbEnv.DB_PARSETIME,
	)

	log.Info(dbSetting)

	db, err := gorm.Open(mysql.Open(dbSetting), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}

	gormAdapter := new(GormAdapter)

	gormAdapter.DB = db

	return gormAdapter
}
