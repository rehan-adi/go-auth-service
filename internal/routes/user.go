package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rehan-adi/go-auth-service/internal/handlers"
	"github.com/rehan-adi/go-auth-service/internal/middlewares"
)

func UserRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	{
		user.GET("/", handlers.GetAllUsers)
		user.GET("/:id", handlers.GetUserById)
		user.PUT("/update", middlewares.AuthMiddleware(), handlers.UpdateUser)
		user.DELETE("/delete", middlewares.AuthMiddleware(), handlers.DeleteUser)
	}
}
