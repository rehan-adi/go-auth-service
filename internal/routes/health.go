package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rehan-adi/go-auth-service/internal/handlers"
)

func HealthRouter(router *gin.Engine) {
	router.GET("/api/v1/health", handlers.HealthCheck)
}
