package config

import "time"

// redis struct
type Redis struct {
	User         string        `yaml:"redis.username"`
	Password     string        `yaml:"redis.password"`
	DB           int           `yaml:"redis.db"`
	Host         string        `yaml:"redis.host"`
	Logger       bool          `yaml:"redis.logger"`
	Port         string        `yaml:"redis.port"`
	UserDuration time.Duration `yaml:"redis.userDuration"`
}
