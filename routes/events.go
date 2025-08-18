package routes

import (
	"events_app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not get events"})
		return
	}
	ctx.JSON(200, events)
}

func getEvent(ctx *gin.Context) {
	idInt64, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid event id"})
		return
	}
	event, err := models.GetEventById(idInt64)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not get event"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Event fetched successfully", "event": event})
}

func updateEvent(ctx *gin.Context) {
	idInt64, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid event id"})
		return
	}
	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(idInt64)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not get event"})
		return
	}
	if event.UserID != userId {
		ctx.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Could not parse request body"})
		return
	}
	updatedEvent.Id = idInt64
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not update event"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Event updated successfully"})
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	userId := ctx.GetInt64("userId")
	event.UserID = userId
	err = event.Save()
	if err != nil {

		ctx.JSON(500, gin.H{"message": "Could not create event"})
		return
	}

	ctx.JSON(201, gin.H{"message": "Event created successfully", "event": event})
}

func deleteEvent(ctx *gin.Context) {
	idInt64, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid event id"})
		return
	}
	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(idInt64)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not get event"})
		return
	}
	if event.UserID != userId {
		ctx.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}
	err = event.Delete()
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not delete event"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Event deleted successfully"})
}
