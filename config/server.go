package config

import "net"

func (g *GRPC) Addr() string {
	return net.JoinHostPort(g.Host, g.Port)
}

func (r *Rest) Addr() string {
	return net.JoinHostPort(r.Host, r.Port)
}
