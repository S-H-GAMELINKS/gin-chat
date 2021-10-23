package controller

import (
	"log"

	"github.com/S-H-GAMELINKS/gin-chat/model"
	"github.com/S-H-GAMELINKS/gin-chat/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type roomController struct {
	db             *gorm.DB
	roomRepository repository.RoomRepository
}

type RoomController interface {
	Index(c *gin.Context)
}

func NewRoomController(db *gorm.DB, roomRepository repository.RoomRepository) RoomController {
	return &roomController{db, roomRepository}
}

func (roomController *roomController) Index(c *gin.Context) {
	var rooms []model.Room

	err := roomController.roomRepository.FindAll(roomController.db, &rooms)
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(200, "rooms/index.tmpl", gin.H{"rooms": rooms})
}
