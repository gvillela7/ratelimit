package model

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"time"
)

type RateLimit struct {
	Client     *redis.Client
	Limit      int
	Expiration time.Duration
	Context    context.Context
}
