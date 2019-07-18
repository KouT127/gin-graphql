package main

import (
	"github.com/KouT127/gin-sample/backend/database"
	"github.com/KouT127/gin-sample/backend/server"
)

func main() {
	database.Init()
	server.Init()
}
