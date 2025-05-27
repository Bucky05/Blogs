package controllers

import (
	"backend/constants"
	"backend/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// db := services.GetDB()
	// db.Query()
}
func Register(c *gin.Context) {
	db := services.GetDB()
	var body map[string]interface{}
	err := c.BindJSON(&body)
	if err != nil {
		return
	}

	username := body["username"]
	email := body["email"]
	password := body["password"]
	name := body["name"]

	msg, err := db.Exec(constants.QueryMap["createUser"], username, email, password, name)
	fmt.Println(msg)

}

// func AuthMiddleware(c *gin.Context) gin.HandlerFunc {
// 	return c.Next()
// }
