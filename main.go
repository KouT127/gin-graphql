package main

import (
	"go-rest/db"
	"go-rest/server"
)

func main() {
	db.Init()
	server.Init()
}
