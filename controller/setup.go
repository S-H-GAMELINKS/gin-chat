package controller

import (
	"github.com/S-H-GAMELINKS/gin-chat/repository"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
)

type interactor struct {
	db *gorm.DB
}

type Interactor interface {
	NewHelloControllerInstance(conn *gorm.DB) HelloController
	NewRoomControllerInstance(conn *gorm.DB) RoomController
	NewMessageControllerInstance(conn *gorm.DB, m *melody.Melody) MessageController
}

func NewInteractor(conn *gorm.DB) Interactor {
	return &interactor{conn}
}

func (interactor *interactor) NewHelloControllerInstance(conn *gorm.DB) HelloController {
	return NewHelloController(conn)
}

func (interactor *interactor) NewRoomControllerInstance(conn *gorm.DB) RoomController {
	return NewRoomController(conn, repository.NewRoomRepositoryInstance())
}

func (interactor *interactor) NewMessageControllerInstance(conn *gorm.DB, m *melody.Melody) MessageController {
	return NewMessageController(conn, m, repository.NewMessageRepositoryInstance())
}
