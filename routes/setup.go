package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(router *gin.Engine, database *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello world!")
	})
}
