package routes

import (
	"events_app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not get event"})
		return
	}
	err = event.Register(userId)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not register for event"})
		return
	}
	ctx.JSON(201, gin.H{"message": "Registered successfully"})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid event id"})
		return
	}
	var event models.Event
	event.Id = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not cancel registration"})
	}
	ctx.JSON(200, gin.H{"message": "Registration cancelled successfully"})
}
