package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rehan-adi/go-auth-service/internal/handlers"
)

func HealthRouter(router *gin.RouterGroup) {
	health := router.Group("/health")
	{
		health.GET("/", handlers.HealthCheck)
	}
}
