package main

import (
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
	"github.com/KouT127/gin-sample/backend/server"
)

func main() {
	database.Init()
	r := server.NewRouter()
	r.Run()
}
