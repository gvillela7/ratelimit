package redis

import (
	"fmt"
	"github.com/gvillela7/ratelimit/configs"
	"github.com/gvillela7/ratelimit/internal/model"
	"golang.org/x/net/context"
	"time"

	"github.com/redis/go-redis/v9"
)

type IRedis interface {
	NewRateLimit(limit int, expiration time.Duration) (*model.RateLimit, error)
	Allow(key string, rl *model.RateLimit) (bool, error)
}

type Redis struct {
	rdb *redis.Client
}

func NewRedis() (IRedis, error) {
	cfg := configs.GetRedisConfig()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	//defer client.Close()
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &Redis{
		rdb: client,
	}, nil
}

func (r Redis) Allow(key string, rl *model.RateLimit) (bool, error) {
	txPipe := rl.Client.TxPipeline()
	increment := txPipe.Incr(rl.Context, key)
	txPipe.Expire(rl.Context, key, rl.Expiration)
	_, err := txPipe.Exec(rl.Context)
	if err != nil {
		return false, err
	}
	return increment.Val() <= int64(rl.Limit), nil
}
func (r Redis) NewRateLimit(limit int, expiration time.Duration) (*model.RateLimit, error) {
	return &model.RateLimit{
		Client:     r.rdb,
		Limit:      limit,
		Expiration: expiration,
		Context:    context.Background(),
	}, nil
}
