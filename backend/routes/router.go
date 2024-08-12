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
		//Category路由
	}

	r.Run(config.GlobalConfig.Server.HttpPort)
}
