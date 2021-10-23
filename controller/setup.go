package controller

import (
	"github.com/S-H-GAMELINKS/gin-chat/repository"
	"gorm.io/gorm"
)

type interactor struct {
	db *gorm.DB
}

type Interactor interface {
	NewHelloControllerInstance(conn *gorm.DB) HelloController
	NewRoomControllerInstance(conn *gorm.DB) RoomController
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
