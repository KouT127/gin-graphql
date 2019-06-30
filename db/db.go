package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-rest/entities"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "go_tutorial"
	CHARSET := "charset=utf8&parseTime=true"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + CHARSET
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	autoMigration()
}

func autoMigration() {
	db.AutoMigrate(&entities.User{})
}

func GetDB() *gorm.DB {
	return db
}
