package controller

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/S-H-GAMELINKS/gin-chat/model"
	"github.com/S-H-GAMELINKS/gin-chat/repository"
	request "github.com/S-H-GAMELINKS/gin-chat/request/message"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
)

type messageController struct {
	db                *gorm.DB
	m                 *melody.Melody
	messageRepository repository.MessageRepository
}

type MessageController interface {
	Create(c *gin.Context)
	Websocket(c *gin.Context)
	Broadcast(s *melody.Session, msg []byte)
}

func NewMessageController(db *gorm.DB, m *melody.Melody, messageRepository repository.MessageRepository) MessageController {
	return &messageController{db, m, messageRepository}
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

func (messageController *messageController) Websocket(c *gin.Context) {
	messageController.m.HandleRequest(c.Writer, c.Request)
}

func (messageController *messageController) Broadcast(s *melody.Session, msg []byte) {

	roomPath := strings.Trim(s.Request.URL.Path, "/roomsmessagews")
	roomID, _ := strconv.ParseUint(roomPath, 10, 64)
	var message model.Message
	message.Content = fmt.Sprintf("%s", msg)
	message.RoomID = roomID
	messageController.messageRepository.Store(messageController.db, &message)

	messageController.m.BroadcastFilter(msg, func(q *melody.Session) bool {
		return q.Request.URL.Path == s.Request.URL.Path
	})
}
