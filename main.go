package main

import (
	"events_app/db"
	"events_app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	err := server.Run(":3000") // localhost:3000
	if err != nil {
		return
	}
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not get events"})
		return
	}
	ctx.JSON(200, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	event.Id = 1
	event.UserID = 2
	err = event.Save()
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not create event"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Event created successfully", "event": event})
}
