package scaffold

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"simcart/app/middleware"
	"simcart/pkg/types"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"simcart/interface/apm"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/logrusorgru/aurora"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (s *skeleton) server(ctx context.Context, suggestedHandlers ...types.RegisterRPCInterface) error {

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	var err error

	s.listener, err = net.Listen("tcp",
		net.JoinHostPort(s.params.ServerGRPC().Host, s.params.ServerGRPC().Port))

	if err != nil {
		return fmt.Errorf("cannot open grpc %v\n", aurora.Red(err))
	}

	grpcServer := grpc.NewServer(grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			otgrpc.OpenTracingStreamServerInterceptor(s.tracer),
			grpc_auth.StreamServerInterceptor(middleware.AuthenticateToken),
		),
	),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_recovery.UnaryServerInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
				otgrpc.OpenTracingServerInterceptor(s.tracer),
				grpc_auth.UnaryServerInterceptor(middleware.AuthenticateToken),
			),
		),
	)

	reflection.Register(grpcServer)

	// register your services here
	conn, err := grpc.DialContext(
		ctx,
		s.params.ServerGRPC().Addr(),
		grpc.WithInsecure(),
	)

	if err != nil {
		return err
	}

	for _, handler := range suggestedHandlers {
		handler(grpcServer, conn, gwmux)
	}

	grpc_prometheus.Register(grpcServer)
	// opentracing
	mux.HandleFunc("/v1/", gwmux.ServeHTTP)
	s.HandleFuncs(mux)

	go grpcServer.Serve(s.listener)
	if err := http.ListenAndServe(
		s.params.ServerRest().Addr(),
		mux); err != nil {

		return err
	}
	return nil

}

// HandleFuncs method for handler your basic methods
func (s *skeleton) HandleFuncs(mux *http.ServeMux) {
	mux.HandleFunc("/v1/apm/metrics", apm.M.Metrics)
	mux.Handle("/v1/apm/health", promhttp.Handler())
}
