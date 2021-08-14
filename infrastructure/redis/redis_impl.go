package redis

import (
	"context"
	"encoding/json"
	"net"
	"simcart/config"
	"simcart/pkg/logger"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// Connect, method for connect to redis
func (r *rds) Connect(confs config.AppConfig) error {
	var err error

	once.Do(func() {

		// logger = zapLogger.GetZapLogger(confs.Mode())

		r.db = redis.NewClient(&redis.Options{
			DB:       r.id,
			Addr:     net.JoinHostPort(confs.ClientRedis().Host, confs.ClientRedis().Port),
			Username: confs.ClientRedis().User,
			Password: confs.ClientRedis().Password,
		})

		if err = r.db.Ping(context.Background()).Err(); err != nil {
			lg := logger.NewPrototype()

			lg.Development().
				Level(zap.ErrorLevel).
				Commit(err.Error())
		}
	})

	return err
}

func (r *rds) setClient(client *redis.Client) {
	r.db = client
}

// Set meth a new key,value
func (r *rds) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	p, err := json.Marshal(value)
	if err != nil {
		lg.
			Development().
			Level(zap.ErrorLevel).
			Commit(err.Error())
		return err
	}
	return r.db.Set(ctx, key, p, duration).Err()
}

// Get meth, get value with key
func (r *rds) Get(ctx context.Context, key string, dest interface{}) error {
	p, err := r.db.Get(ctx, key).Result()

	if err != nil {
		lg.
			Development().
			Level(zap.ErrorLevel).
			Commit(err.Error())
		return err
	}

	return json.Unmarshal([]byte(p), &dest)
}

// Del for delete keys in redis
func (r *rds) Del(ctx context.Context, key ...string) error {
	_, err := r.db.Del(ctx, key...).Result()
	if err != nil {
		lg.
			Development().
			Level(zap.ErrorLevel).
			Commit(err.Error())
		return err
	}
	return nil

}
