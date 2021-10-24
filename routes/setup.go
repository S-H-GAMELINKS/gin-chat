package routes

import (
	"log"

	"github.com/S-H-GAMELINKS/gin-chat/controller"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
)

func SetUpRoutes(router *gin.Engine, database *gorm.DB) {
	router.LoadHTMLGlob("template/**/*")

	m := melody.New()

	interactor := controller.NewInteractor(database)

	helloController := interactor.NewHelloControllerInstance(database)

	router.GET("/", helloController.Index)

	roomController := interactor.NewRoomControllerInstance(database)

	router.GET("/rooms", roomController.Index)
	router.GET("/rooms/:id", roomController.Show)
	router.POST("/rooms", roomController.Create)

	messageController := interactor.NewMessageControllerInstance(database, m)

	router.POST("/rooms/:id/messages", messageController.Create)
	router.GET("/rooms/:id/messages/ws", messageController.Websocket)

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	m.HandleConnect(func(s *melody.Session) {
		log.Printf("websocket connection open. [session: %#v]\n", s)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		log.Printf("websocket connection close. [session: %#v]\n", s)
	})

	m.HandleMessage(messageController.Broadcast)
}
