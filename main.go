package main

import (
	"events_app/db"
	"events_app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	err := godotenv.Load()
	if err != nil {
		return
	}

	err = server.Run(":3000") // localhost:3000
	if err != nil {
		return
	}
}
