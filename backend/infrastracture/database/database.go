package database

import (
	"gin-sample/backend/domain/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	OPTION := "charset=utf8&parseTime=true"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	autoMigration()
}

func autoMigration() {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Task{})
}

func GetDB() *gorm.DB {
	return db
}
