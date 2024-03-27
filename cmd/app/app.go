package main

import (
	"fmt"

	"github.com/cyborgvova/report/database"
	"github.com/cyborgvova/report/server"
)

func main() {
	// report := &database.Report{}
	database.CreateTable()
	// db := database.DBConnect()
	// db.Find(&report)
	// fmt.Println(report)
	fmt.Println("Hello World !!!")
	server.NewServer()
}
