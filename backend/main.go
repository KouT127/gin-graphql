package main

import (
	"go-rest/backend/db"
	"go-rest/backend/server"
)

func main() {
	db.Init()
	server.Init()
}
