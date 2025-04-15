package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginmiddleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimiterMiddleware() gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted("100-M")
	if err != nil {
		panic(err)
	}

	store := memory.NewStore()
	instance := limiter.New(store, rate)

	return ginmiddleware.NewMiddleware(instance)
}
