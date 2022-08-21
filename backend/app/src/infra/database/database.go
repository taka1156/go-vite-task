package database

import (
	"app/envars"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SqlAdapter struct {
	DB *sql.DB
}

func isErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func DbInit() *SqlAdapter {
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
	db, err := sql.Open(envars.DbEnv.DB_NAME, dbSetting)
	if err != nil {
		panic(err.Error)
	}

	sqlAdapter := new(SqlAdapter)

	sqlAdapter.DB = db

	return sqlAdapter
}
