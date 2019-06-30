package main

import (
	"gin-sample/backend/database"
	"gin-sample/backend/server"
)

func main() {
	database.Init()
	server.Init()
}
