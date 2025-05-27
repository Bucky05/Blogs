package main

import (
	"backend/database"
	"backend/routes"
	"backend/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()
	services.InitDB(db)
	defer db.Close()
	if err != nil {
		log.Fatal("DB connection failed: ", err)
	}
	fmt.Println("DB connection successful")
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")

	defer db.Close()
}
