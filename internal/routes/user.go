package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rehan-adi/go-auth-service/internal/handlers"
)

func UserRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		user.GET("/", handlers.GetAllUsers)
	}
}
