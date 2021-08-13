package scaffold

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"simcart/api/pb/cartpb"
	"simcart/api/pb/productpb"
	"simcart/api/pb/productpb/searchpb"
	"simcart/app/middleware"
	"simcart/config"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"simcart/interface/apm"
	grpc_cart_handler "simcart/interface/product"
	grpc_product_handler "simcart/interface/product"
	grpc_product_search_handler "simcart/interface/product/search"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/logrusorgru/aurora"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/reflection"

	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

func (s *skeleton) server(ctx context.Context) error {
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
				otgrpc.OpenTracingServerInterceptor(s.tracer, otgrpc.LogPayloads()),
				grpc_auth.UnaryServerInterceptor(middleware.AuthenticateToken),
			),
		),
	)
	go func() {
		// add reflection in debug mode
		if config.GetAppConfig().Mode() {
			reflection.Register(grpcServer)
		}
		grpcServer.Serve(s.listener)

	}()

	// register your services here
	// we pass main context to all the underlying servers to implement
	// graceful shutdown if needed
	productpb.RegisterProductServiceServer(grpcServer,
		grpc_product_handler.NewProductServerHandler(ctx))

	searchpb.RegisterSearchServer(grpcServer,
		grpc_product_search_handler.NewSearchServerHandler(ctx))

	cartpb.RegisterCartServer(grpcServer, grpc_cart_handler.NewProductServerHandler(ctx))

	healthgrpc.RegisterHealthServer(grpcServer, health.NewServer())

	// dial context
	conn, err := grpc.DialContext(
		context.Background(),
		s.params.ServerGRPC().Addr(),
		grpc.WithInsecure(),
	)

	if err != nil {
		return err
	}

	if err = productpb.RegisterProductServiceHandler(ctx, gwmux, conn); err != nil {
		log.Fatalf("Failed to connect to register gateway: %v\n", err)
	}

	if err = searchpb.RegisterSearchHandler(ctx, gwmux, conn); err != nil {
		log.Fatalf("Failed to connect to register gateway: %v\n", err)
	}

	if err := cartpb.RegisterCartHandler(ctx, gwmux, conn); err != nil {
		log.Fatalf("Failed to connect to register gateway: %v\n", err)
	}

	// prometheus tracing
	grpc_prometheus.Register(grpcServer)

	// do tracing and prepare metrics

	mux.HandleFunc("/v1/", gwmux.ServeHTTP)
	s.HandleFuncs(mux)

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
