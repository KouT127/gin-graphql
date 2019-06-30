package main

import (
	"gin-sample/backend/db"
	"gin-sample/backend/server"
)

func main() {
	db.Init()
	server.Init()
}
