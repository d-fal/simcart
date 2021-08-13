package config

import "net"

// redis struct
type RediSearch struct {
	User     string `yaml:"redisearch.username"`
	Password string `yaml:"redisearch.password"`
	Host     string `yaml:"redisearch.host"`
	Port     string `yaml:"redisearch.port"`
}

func (r *RediSearch) Addr() string {
	return net.JoinHostPort(r.Host, r.Port)
}
