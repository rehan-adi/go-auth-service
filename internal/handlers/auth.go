package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rehan-adi/go-auth-service/internal/database"
	"github.com/rehan-adi/go-auth-service/internal/models"
	"github.com/rehan-adi/go-auth-service/internal/utils"
	"github.com/rehan-adi/go-auth-service/internal/validators"
)

func Signup(ctx *gin.Context) {

	var data validators.SignupValidator

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	errors := validators.ValidateSignupData(data)

	if len(errors) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"validation errors": errors})
		return
	}

	var existingUser models.User
	if err := database.DB.Where("email = ?", data.Email).First(&existingUser).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	hashpassword, err := utils.HashPassword(data.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	user := models.User{
		Email:    data.Email,
		Username: data.Username,
		Password: string(hashpassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

}

func Signin(ctx *gin.Context) {

	var data validators.SigninValidator

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request format"})
		return
	}

	validationErrors := validators.ValidateSigninData(data)

	if len(validationErrors) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "validation_error": validationErrors})
		return
	}

	var user models.User

	if err := database.DB.Where("email = ?", data.Email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "error": "User not found"})
		return
	}

	isValid := utils.VerifyPassword(data.Password, user.Password)

	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Could not generate token"})
		return
	}

	ctx.SetCookie("token", token, 86400, "/", "", true, false)

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": token, "message": "Login successful"})

}
