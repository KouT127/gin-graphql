package database

import (
	"github.com/KouT127/gin-sample/backend/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

func Init(c *config.Config) {
	DBMS := c.Database.Dbms
	USER := c.Database.User
	PASS := c.Database.Pass
	PROTOCOL := "tcp(" + c.Database.Host + ":" + c.Database.Port + ")"
	DBNAME := c.Database.DbName
	OPTION := c.Database.Option

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
}

func NewDB() *gorm.DB {
	return db
}
