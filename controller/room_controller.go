package controller

import (
	"log"
	"strconv"

	"github.com/S-H-GAMELINKS/gin-chat/model"
	"github.com/S-H-GAMELINKS/gin-chat/repository"
	request "github.com/S-H-GAMELINKS/gin-chat/request/room"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type roomController struct {
	db             *gorm.DB
	roomRepository repository.RoomRepository
}

type RoomController interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Create(c *gin.Context)
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

func (roomController *roomController) Show(c *gin.Context) {

	roomID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var room model.Room
	room.ID = roomID

	err := roomController.roomRepository.Find(roomController.db, &room)
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(200, "rooms/show.tmpl", gin.H{"room": room})
}

func (roomController *roomController) Create(c *gin.Context) {

	var createRoomRequest request.CreateRoomRequest

	err := c.Bind(&createRoomRequest)
	if err != nil {
		log.Fatal(err)
	}

	var room model.Room

	room.Name = createRoomRequest.Name

	err = roomController.roomRepository.Store(roomController.db, &room)
	if err != nil {
		log.Fatal(err)
	}

	redirectPath := "/rooms/" + strconv.FormatUint(room.ID, 10)

	c.Redirect(302, redirectPath)
}
