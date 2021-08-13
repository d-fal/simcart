package redis

import (
	"context"
	"simcart/config"
	"simcart/pkg/logger"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	lg   *logger.Log
	once sync.Once
)

// store interface is interface for store things into redis
type store interface {
	Connect(config.AppConfig) error
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string, dest interface{}) error
	Del(ctx context.Context, key ...string) error
	setClient(client *redis.Client)
}

// rds struct for redis client
type rds struct {
	db *redis.Client
	id int
}

func NewClient(db int) store {
	return &rds{id: db}
}
