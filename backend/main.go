package main

import (
	"github.com/KouT127/gin-sample/backend/config"
	"github.com/KouT127/gin-sample/backend/infrastracture"
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
)

func main() {
	config.Init(config.Development)
	c := config.NewConfig()
	database.Init(c)
	infrastracture.Init()
}
