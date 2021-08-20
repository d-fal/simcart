package scaffold

import (
	"context"
	"fmt"
	"html"
	"log"
	"simcart/config"
	"simcart/infrastructure/postgres"
	"simcart/infrastructure/search"

	"github.com/logrusorgru/aurora"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"github.com/uber/jaeger-lib/metrics/prometheus"
)

type RunningOption func(c context.Context, s *skeleton) error

func WithPostgres() RunningOption {
	return func(c context.Context, s *skeleton) error {
		logLevel := config.Silent

		if s.params.Mode() {
			logLevel = config.Verbose
		}
		if err := postgres.Storage.Connect(logLevel, config.GetAppConfig().ClientPostgres()); err != nil {
			return fmt.Errorf("%v \t cannot initialize postgresql at: %s | %v\n", config.GetAppConfig().ClientPostgres().Host, aurora.White(html.UnescapeString("&#x274C;")), err)
		}

		return nil
	}
}

func WithOpenTracing() RunningOption {
	return func(c context.Context, s *skeleton) error {
		var err error

		metricsFactory := prometheus.New()

		cfg := jaegercfg.Configuration{
			ServiceName: config.GetAppConfig().Application().Id,

			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans:           true,
				LocalAgentHostPort: config.GetAppConfig().ClientJaeger().Addr(),
			},
		}

		jLogger := jaegerlog.StdLogger
		jMetricsFactory := metrics.NullFactory

		// Initialize tracer with a logger and a metrics factory
		s.tracer, s.closer, err = cfg.NewTracer(
			jaegercfg.Logger(jLogger),
			jaegercfg.Metrics(jMetricsFactory),
			jaegercfg.Metrics(metricsFactory),
		)
		if err != nil {
			log.Printf("Could not initialize jaeger tracer: %s", err.Error())

		}

		opentracing.SetGlobalTracer(s.tracer)

		if err != nil {
			return fmt.Errorf("%v \t cannot initialize jaeger  %v\n", aurora.White(html.UnescapeString("&#x274C;")), err)
		}
		return nil
	}
}

func WithRedisearch() RunningOption {
	return func(c context.Context, s *skeleton) error {
		client := search.NewClient(config.GetAppConfig().ClientRedisearch().Addr(), "indexer")

		if client == nil {
			return fmt.Errorf("%v initializing redisearch client failed.\n",
				aurora.Red(html.UnescapeString("&#x274C;")))
		}
		return nil
	}
}
