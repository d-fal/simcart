package scaffold

import (
	"context"
	"io"
	"net"
	"net/http"
	"simcart/config"

	"github.com/opentracing/opentracing-go"
)

type Scaffold interface {
	server(context.Context) error
	tracing() (opentracing.Tracer, io.Closer, error)

	commissioning(ctx context.Context) error

	Close()

	SetConfigPath(string)

	// for testing purposes
	Hibernate(ctx context.Context) error

	NormalStart(ctx context.Context) error

	HandleFuncs(mux *http.ServeMux)
}

type skeleton struct {
	params   config.AppConfig
	listener net.Listener
	tracer   opentracing.Tracer
	closer   io.Closer
	context  context.Context
	cancel   context.CancelFunc
}
