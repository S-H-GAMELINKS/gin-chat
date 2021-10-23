package repository

import (
	"log"

	"github.com/S-H-GAMELINKS/gin-chat/model"
	"gorm.io/gorm"
)

type messageRepository struct {
}

type MessageRepository interface {
	Store(db *gorm.DB, message *model.Message) (err error)
}

func NewMessageRepository() MessageRepository {
	return &messageRepository{}
}

func (messageRepository *messageRepository) Store(db *gorm.DB, message *model.Message) (err error) {
	if err = db.Create(&message).Error; err != nil {
		log.Fatal(err)
	}
	return
}
