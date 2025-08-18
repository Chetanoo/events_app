package main

import (
	"events_app/db"
	"events_app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":3000") // localhost:3000
	if err != nil {
		return
	}
}
