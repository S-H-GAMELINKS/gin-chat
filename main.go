package main

import (
	"github.com/S-H-GAMELINKS/gin-chat/database"
	"github.com/S-H-GAMELINKS/gin-chat/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	conn, err := database.SetUpDB()
	if err != nil {
		panic(err)
	}

	routes.SetUpRoutes(r, conn)

	r.Run(":8000")
}
