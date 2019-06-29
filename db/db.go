package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "go_tutorial"
	CHARSET := "charset=utf8"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + CHARSET
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
