package main

import (
	"fmt"

	"github.com/jonatasemanuel/go-sqlx/db"
)

func main() {
	fmt.Println("Application start up")
	_, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Error creating database")
	}

	fmt.Println("successfully connected to database")
	return
}
