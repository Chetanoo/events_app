package routes

import (
	"events_app/models"
	"events_app/utils"

	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
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

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Could not parse user data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(401, gin.H{"message": "Invalid credentials"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Could not generate token"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User authenticated successfully", "token": token})
}
