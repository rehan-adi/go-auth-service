package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rehan-adi/go-auth-service/internal/database"
	"github.com/rehan-adi/go-auth-service/internal/dto"
	"github.com/rehan-adi/go-auth-service/internal/models"
	"github.com/rehan-adi/go-auth-service/internal/utils"
	"github.com/rehan-adi/go-auth-service/internal/validators"
)

func GetAllUsers(ctx *gin.Context) {

	var users []models.User
	var usersResponse []dto.UserDataResponse

	if err := database.DB.Select("id, username, email, created_at").
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
			Created:  user.CreatedAt.Format("2006-01-02 15:04:05"),
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

	if err := database.DB.Select("id, email, username, created_at").Where("id = ?", id).
		First(&user).Error; err != nil {
		utils.Log.Errorf("User not found: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "User not found",
		})
		return
	}

	fmt.Println(user)

	usersResponse := dto.UserDataResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Created:  user.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usersResponse,
		"message": "User retrieved successfully",
	})

}

func UpdateUser(ctx *gin.Context) {

	id, exists := ctx.Get("id")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Unauthorized: User ID not found",
		})
		return
	}

	var data validators.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid input: Name is required (3-50 characters)",
		})
		return
	}

	validationErrors := validators.ValidateUpdateUserData(data)
	if len(validationErrors) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errors":  validationErrors,
		})
		return
	}

	var user models.User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "User not found",
		})
		return
	}

	result := database.DB.Model(&user).Update("username", data.Username)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to update user details",
		})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "No changes made to user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User name updated successfully",
	})

}

func DeleteUser(ctx *gin.Context) {

	id, exists := ctx.Get("id")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Unauthorized: User ID not found",
		})
		return
	}

	result := database.DB.Delete(&models.User{}, "id = ?", id)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to delete user",
		})
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User deleted successfully",
	})

}
