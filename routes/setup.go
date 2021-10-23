package routes

import (
	"github.com/S-H-GAMELINKS/gin-chat/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(router *gin.Engine, database *gorm.DB) {
	router.LoadHTMLGlob("template/**/*")

	interactor := controller.NewInteractor(database)

	helloController := interactor.NewHelloControllerInstance(database)

	router.GET("/", helloController.Index)

	roomController := interactor.NewRoomControllerInstance(database)

	router.GET("/rooms", roomController.Index)
}
