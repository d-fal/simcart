package config

import "net"

// Jaeger tracer
type Jaeger struct {
	Host string
	Port string
}

func (j *Jaeger) Addr() string {
	return net.JoinHostPort(j.Host, j.Port)
}
