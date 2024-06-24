package main

import (
	"e-commerce-back/pkg/db"
	"fmt"
)

func main() {
	database := db.DBConnection()

	fmt.Println("Running migrations")

	database.AutoMigrate()
}
