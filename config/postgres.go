package config

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type LogLevel int

const (
	Silent LogLevel = (iota) << 1
	Verbose
)
