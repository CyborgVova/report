package main

import (
	"github.com/cyborgvova/report/database"
	"github.com/cyborgvova/report/server"
)

func main() {
	database.CreateTable()
	server.NewServer()
}
