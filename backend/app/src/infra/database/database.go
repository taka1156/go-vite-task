package database

import (
	"app/infra/envars"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type GormHandler struct {
	DB *gorm.DB
}

func isErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func DbInit() *GormHandler {
	// "user:password@tcp(db:3306)/sample?"
	dbSetting := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&collation=%s&parseTime=%s",
		envars.DbEnv.DB_USER,
		envars.DbEnv.DB_PASSWORD,
		envars.DbEnv.DB_PORT,
		envars.DbEnv.DB_DATABASE,
		envars.DbEnv.CHARSET,
		envars.DbEnv.COLLATION,
		envars.DbEnv.PARSETIME,
	)
	db, err := gorm.Open(envars.DbEnv.DB_NAME, dbSetting)
	if err != nil {
		panic(err.Error)
	}

	db.LogMode(true)

	gormHandler := new(GormHandler)

	gormHandler.DB = db

	return gormHandler
}
