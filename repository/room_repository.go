package repository

import (
	"log"

	"github.com/S-H-GAMELINKS/gin-chat/model"
	"gorm.io/gorm"
)

type roomRepository struct {
}

type RoomRepository interface {
	FindAll(db *gorm.DB, rooms *[]model.Room) (err error)
	Store(db *gorm.DB, room *model.Room) (err error)
}

func NewRoomRepository() RoomRepository {
	return &roomRepository{}
}

func (roomRepository *roomRepository) FindAll(db *gorm.DB, rooms *[]model.Room) (err error) {
	if err = db.Find(&rooms).Error; err != nil {
		log.Fatal(err)
	}
	return
}

func (roomRepository *roomRepository) Store(db *gorm.DB, room *model.Room) (err error) {
	if err = db.Create(&room).Error; err != nil {
		log.Fatal(err)
	}
	return
}
