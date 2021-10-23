package controller

import (
	"log"
	"strconv"

	"github.com/S-H-GAMELINKS/gin-chat/model"
	"github.com/S-H-GAMELINKS/gin-chat/repository"
	request "github.com/S-H-GAMELINKS/gin-chat/request/message"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type messageController struct {
	db                *gorm.DB
	messageRepository repository.MessageRepository
}

type MessageController interface {
	Create(c *gin.Context)
}

func NewMessageController(db *gorm.DB, messageRepository repository.MessageRepository) MessageController {
	return &messageController{db, messageRepository}
}

func (messageController *messageController) Create(c *gin.Context) {
	var createMessageRequest request.CreateMessageRequest

	err := c.Bind(&createMessageRequest)
	if err != nil {
		log.Fatal(err)
	}

	roomID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var message model.Message
	message.Content = createMessageRequest.Content
	message.RoomID = roomID

	err = messageController.messageRepository.Store(messageController.db, &message)
	if err != nil {
		log.Fatal(err)
	}
	redirectPath := "/rooms/" + c.Param("id")
	c.Redirect(302, redirectPath)
}
