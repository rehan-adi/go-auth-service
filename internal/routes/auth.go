package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rehan-adi/go-auth-service/internal/handlers"
)

func AuthRouter(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", handlers.Signup)
		auth.POST("/signin", handlers.Signin)
	}
}
