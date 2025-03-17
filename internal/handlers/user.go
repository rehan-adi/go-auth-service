package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rehan-adi/go-auth-service/internal/database"
	"github.com/rehan-adi/go-auth-service/internal/dto"
	"github.com/rehan-adi/go-auth-service/internal/models"
	"github.com/rehan-adi/go-auth-service/internal/utils"
)

func GetAllUsers(ctx *gin.Context) {

	var users []models.User
	var usersResponse []dto.UserDataResponse

	if err := database.DB.Select("id, username, email, created").
		Find(&users).Error; err != nil {
		utils.Log.Errorf("Failed to retrieve users: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to retrieve users",
		})
		return
	}

	for _, user := range users {
		usersResponse = append(usersResponse, dto.UserDataResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			Created:  user.Created.Format("2006-01-02 15:04:05"),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usersResponse,
		"message": "All users retrieved successfully",
	})

}

func GetUserById(ctx *gin.Context) {

	var user models.User

	id := ctx.Param("id")

	if err := database.DB.Select("id, email, username, created").Where("id = ?", id).
		First(&user).Error; err != nil {
		utils.Log.Errorf("User not found: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "User not found",
		})
		return
	}

	usersResponse := dto.UserDataResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Created:  user.Created.Format("2006-01-02 15:04:05"),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usersResponse,
		"message": "User retrieved successfully",
	})

}
