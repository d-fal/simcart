package config

type AppConfig interface {
	load() error
	Debug(bool)

	// set mode, debug or release
	Mode() bool
	// set log path
	SetPath(string)
	Level(uint8)

	Application() *App

	ServerGRPC() *GRPC
	ServerRest() *Rest

	ClientRedis() *Redis
	ClientJaeger() *Jaeger
	ClientPostgres() *Postgres
	ClientRedisearch() *RediSearch
}

type appConfig struct {
	debug bool
	base  string
	level uint8 // log level

	*App

	*Server

	*Client
}

type Parameters struct{}

var config AppConfig

func init() {
	config = new(appConfig)
}

func GetAppConfig() AppConfig {
	return config
}

type Server struct {
	*GRPC
	*Rest
}

type Client struct {
	*Redis
	*Jaeger
	*RediSearch
	*Postgres
}

type GRPC struct {
	Host string
	Port string
}

type Rest struct {
	Host string
	Port string
}
