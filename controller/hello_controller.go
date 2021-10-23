package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type helloController struct {
	db *gorm.DB
}

type HelloController interface {
	Index(c *gin.Context)
}

func NewHelloController(conn *gorm.DB) HelloController {
	return &helloController{conn}
}

func (helloController *helloController) Index(c *gin.Context) {
	c.HTML(200, "hello/index.tmpl", gin.H{})
}
