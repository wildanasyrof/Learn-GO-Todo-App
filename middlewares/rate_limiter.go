package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

// RateLimitMiddleware applies rate limiting to routes
func RateLimitMiddleware() gin.HandlerFunc {
	// Define limit: 10 requests per minute per IP
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  10, // Change this based on your needs
	}

	// Use in-memory storage
	store := memory.NewStore()

	// Create rate limiter
	rateLimiter := limiter.New(store, rate)

	// Apply middleware
	return ginlimiter.NewMiddleware(rateLimiter)
}
