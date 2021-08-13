package config

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/viper"
)

func (c *appConfig) Debug(debug bool) {
	c.debug = debug
}

func (c *appConfig) Mode() bool {
	return c.debug
}

func (c *appConfig) SetPath(path string) {

	c.base = path

	if err := config.load(); err != nil {
		panic(fmt.Errorf("cannot unmarshal config %s", err))
	}
}

func (c *appConfig) load() error {

	viper.AddConfigPath(c.base)
	viper.SetConfigName(".config")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config file not found: %v\n", aurora.Red(err)))
	}

	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}

func (c *appConfig) Level(level uint8) {
	c.level = level
}

func (c *appConfig) Application() *App {
	return c.App
}

func (c *appConfig) ServerGRPC() *GRPC {
	return c.GRPC
}

func (c *appConfig) ServerRest() *Rest {
	return c.Rest
}

func (c *appConfig) ClientRedis() *Redis {
	return c.Redis
}

func (c *appConfig) ClientJaeger() *Jaeger {
	return c.Jaeger
}

func (c *appConfig) ClientPostgres() *Postgres {
	return c.Postgres
}

func (c *appConfig) ClientRedisearch() *RediSearch {
	return c.RediSearch
}
