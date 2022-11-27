package database

import (
	"app/envvars"
	"database/sql"
	"fmt"

	"github.com/labstack/gommon/log"
)

type SqlAdapter struct {
	DB *sql.DB
}

var SQL = &SqlAdapter{}

func DBInit(DbEnv envvars.DbSetting) *SqlAdapter {
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

	db, err := sql.Open("mysql", dbSetting)
	if err != nil {
		panic(err.Error)
	}

	SQL.DB = db

	return SQL
}
