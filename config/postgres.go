package config

import "net"

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

func (cnf *Postgres) Addr() string {
	return net.JoinHostPort(
		cnf.Host,
		cnf.Port,
	)
}
