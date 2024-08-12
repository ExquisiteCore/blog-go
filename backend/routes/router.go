package routes

import (
	"backend/api"
	"backend/config"

	"github.com/gin-gonic/gin"
)

func InitRuter() {
	gin.SetMode(config.GlobalConfig.Server.AppMode)

	r := gin.Default()

	router := r.Group("/api")
	{
		//User路由
		user := router.Group("/user")
		{
			user.POST("/add", api.AddUser)
			user.DELETE("/delete/:name", api.DeleteUser)
		}
		//Post路由
		Post := router.Group("/post")
		{
			Post.POST("/add", api.AddPost)
			Post.DELETE("/delete/:id", api.DeletePost)
			Post.GET("/get", api.GetPost)
			Post.GET("/get/:id", api.GetPostInfo)
		}
		//Category路由
		Category := router.Group("/category")
		{
			Category.POST("/add", api.AddCategory)
			Category.DELETE("/delete/:name", api.DeleteCategory)
			Category.GET("/get", api.GetCate)
			Category.GET("/get/:name", api.GetCateInfo)
		}
	}

	r.Run(config.GlobalConfig.Server.HttpPort)
}
