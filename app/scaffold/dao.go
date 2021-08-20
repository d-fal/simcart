package scaffold

import (
	"context"
	"io"
	"net"
	"net/http"
	"simcart/config"
	internal_types "simcart/pkg/types"

	"github.com/opentracing/opentracing-go"
)

type Scaffold interface {
	server(context.Context, ...internal_types.RegisterRPCInterface) error

	Close()

	SetConfigPath(string)

	Start(ctx context.Context, runServer bool, options ...RunningOption) internal_types.SelectedInterfaces

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
