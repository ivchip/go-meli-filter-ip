package http

import (
	"github.com/go-redis/redis/v8"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	redisLimiter "github.com/ulule/limiter/v3/drivers/store/redis"
	"log"
	"net/http"
	"os"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

func (m *GoMiddleware) LimitRate(next http.Handler) http.Handler {
	redisAddr := os.Getenv("REDIS_ADDRESS")
	redisPass := os.Getenv("REDIS_PASS")
	limitCmd := os.Getenv("LIMIT_COMMAND")

	// Define a limit rate to number requests per hour.
	rate, err := limiter.NewRateFromFormatted(limitCmd)
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(&redis.Options{Addr: redisAddr, Password: redisPass, DB: 0})

	// Create a store with the redis client.
	store, err := redisLimiter.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix: "limiter_chi",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new middleware with the limiter instance.
	middlewareLimit := stdlib.NewMiddleware(limiter.New(store, rate, limiter.WithTrustForwardHeader(true)))
	return middlewareLimit.Handler(next)
}

// InitMiddleware initialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
