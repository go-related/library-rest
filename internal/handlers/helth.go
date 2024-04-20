package handlers

import "github.com/gin-gonic/gin"

func SetupHealth(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

}
