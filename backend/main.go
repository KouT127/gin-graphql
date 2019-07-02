package main

import (
	"gin-sample/backend/infrastracture/database"
	"gin-sample/backend/infrastracture/server"
)

func main() {
	database.Init()
	server.Init()
}
