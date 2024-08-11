package routes

import (
	"backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRuter() {
	gin.SetMode(config.GlobalConfig.Server.AppMode)

	r := gin.Default()

	router := r.Group("/api")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(config.GlobalConfig.Server.HttpPort)
}
