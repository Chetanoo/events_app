package routes

import (
	"events_app/models"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Could not parse user data"})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not create user"})
		return
	}
	ctx.JSON(201, gin.H{"message": "User created successfully", "user": user})
}
