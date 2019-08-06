package main

import (
	"github.com/KouT127/gin-sample/backend/config"
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
	"github.com/KouT127/gin-sample/backend/server"
)

func main() {
	config.Init()
	database.Init()
	server.Init()
}
