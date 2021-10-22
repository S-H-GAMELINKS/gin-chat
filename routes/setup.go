package routes

import (
	"github.com/S-H-GAMELINKS/gin-chat/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(router *gin.Engine, database *gorm.DB) {
	interactor := controller.NewInteractor(database)

	helloController := interactor.NewHelloControllerInstance(database)

	router.GET("/", helloController.Index)
}
