package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type gormAdapter struct {
	db *gorm.DB
}

func isErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func dbConn() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "user:password@tcp(db:3306)/sample?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true")
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	return db, nil
}
