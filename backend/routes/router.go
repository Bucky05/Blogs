package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	blog := router.Group("/blog")
	//blog.Use(controllers.AuthMiddleware())
	{
		blog.GET("/", controllers.GetAllBlogs)
		blog.GET(":id", controllers.GetBlogByID)
		blog.POST("/", controllers.CreateBlog)
		blog.DELETE(":id", controllers.DeleteBlog)
		blog.PUT(":id", controllers.UpdateBlog)
	}

	// rating := router.Group("/ratings")
	// {
	// 	rating.POST("/", controllers.RateBlog)
	// }

	// comment := router.Group("/comments")
	// {
	// 	comment.POST("/", controllers.CreateComment)
	// 	comment.GET("/blog/:id", controllers.GetCommentsForBlog)
	// }

}
