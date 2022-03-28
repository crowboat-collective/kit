package redis

import (
	"context"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Client holds a reference to a redis Pool.
type Client struct {
	pool *redis.Pool
}

// Connect attempts a connection the redis db defined in the REDIS_CONNECTION_URI environment variable.
func Connect(ctx context.Context) (*Client, error) {
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.DialContext(ctx, "tcp", os.Getenv("REDIS_CONNECTION_URI")) },
	}

	return &Client{
		pool: pool,
	}, nil
}
