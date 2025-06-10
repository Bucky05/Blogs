package controllers

import (
	"backend/queries"
	"backend/services"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Login(c *gin.Context) {
	db := services.GetDB()
	var body map[string]interface{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": " Invalid request body"})
		return
	}
	username := body["username"]
	password := body["password"]
	var pass string
	err = db.QueryRow(queries.GetPasswordByUsername(), username).Scan(&pass)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(401, gin.H{"error": "No User found"})
		} else {
			log.Fatal("Failed to query user: ", err)
			c.JSON(500, gin.H{"error": "Internal server error"})

		}

	} else if pass == password {
		c.JSON(200, gin.H{"message": "Login successful"})
	} else {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
	}
}
func Register(c *gin.Context) {
	db := services.GetDB()
	var body map[string]interface{}
	err := c.BindJSON(&body)
	if err != nil {
		return
	}
	newUUID, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("Failed to generate UUID: ", err)
	}
	uuidString := newUUID.String()
	username := body["username"]
	email := body["email"]
	password := body["password"]
	name := body["name"]

	msg, err := db.Exec(queries.CreateUser(), uuidString, username, email, password, name)
	if err != nil {
		log.Fatal("Failed to create user: ", err)
	}
	fmt.Print(msg)
	c.JSON(201, gin.H{
		"message": fmt.Sprintf("User %s created successfully", username),
	})

}

// func AuthMiddleware(c *gin.Context) gin.HandlerFunc {
// 	return c.Next()
// }
