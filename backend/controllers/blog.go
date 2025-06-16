package controllers

import (
	"backend/services"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	Title      string
	Content    string
	UserID     string
	CreateDate string
	ID         string
	Category   string
}

func CreateBlog(c *gin.Context) {
	var body map[string]interface{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
	}

	blog := Blog{
		ID:      services.GenerateUUID(),
		Title:   body["title"].(string),
		Content: body["content"].(string),
	}
}
func GetAllBlogs(c *gin.Context) {

}
func GetBlogByID(c *gin.Context) {

}
func UpdateBlog(c *gin.Context) {

}
func DeleteBlog(c *gin.Context) {

}
func RateBlog(c *gin.Context) {
}
func CreateComment(c *gin.Context) {
}
