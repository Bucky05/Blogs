package main

import (
	"backend/database"
	"log"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		log.Fatal("DB connection failed: ", err)
	}

	defer db.Close()
}
