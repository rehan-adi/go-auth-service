package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rehan-adi/go-auth-service/internal/database"
	"github.com/rehan-adi/go-auth-service/internal/dto"
	"github.com/rehan-adi/go-auth-service/internal/models"
)

func GetAllUsers(ctx *gin.Context) {

	var users []dto.UserDataResponse

	if err := database.DB.Model(&models.User{}).Select("id, username, email, created").Find(&users).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to retrieve users",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
		"message": "All users retrieved successfully",
	})

}
