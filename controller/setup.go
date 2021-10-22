package controller

import (
	"gorm.io/gorm"
)

type interactor struct {
	db *gorm.DB
}

type Interactor interface {
	NewHelloControllerInstance(conn *gorm.DB) HelloController
}

func NewInteractor(conn *gorm.DB) Interactor {
	return &interactor{conn}
}

func (interactor *interactor) NewHelloControllerInstance(conn *gorm.DB) HelloController {
	return NewHelloController(conn)
}
