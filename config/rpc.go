package config

import "net"

type GRPCClient struct {
	Host string
	Port string
	TLS  *string
	Ca   *string
}

func (g *GRPCClient) TLSEnabled() bool {
	return g.TLS != nil
}

func (g *GRPCClient) Nil() bool {
	return g == nil
}

func (g *GRPCClient) Defined() bool {

	if g.Host != "" && g.Port != "" {
		return true
	}

	return false

}

func (g *GRPCClient) Addr() string {
	return net.JoinHostPort(g.Host, g.Port)
}
