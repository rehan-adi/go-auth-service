package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rehan-adi/go-auth-service/internal/handlers"
)

func AuthRouter(router *gin.Engine) {
	router.POST("/api/v1/auth/signup", handlers.Signup)
	router.POST("/api/v1/auth/signin", handlers.Signin)
}
