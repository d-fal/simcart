package scaffold

import (
	"context"
	"fmt"
	"html"
	"simcart/config"
	"simcart/infrastructure/postgres"
	"simcart/infrastructure/search"
	"simcart/pkg/logger"

	"github.com/logrusorgru/aurora"
	"go.uber.org/zap/zapcore"
)

func (s *skeleton) commissioning(ctx context.Context) (err error) {

	s.context, s.cancel = context.WithCancel(ctx)

	s.tracer, s.closer, err = s.tracing()

	if err != nil {
		logger.NewPrototype().Add("jaeger", err).Level(zapcore.WarnLevel).Commit("jaeger")
	}

	logLevel := config.Silent

	if s.params.Mode() {
		logLevel = config.Verbose
	}

	if err = postgres.Storage.Connect(logLevel,
		config.GetAppConfig().ClientPostgres()); err != nil {
		return fmt.Errorf("%v cannot initializing db %v at: %s\n",
			aurora.White(html.UnescapeString("&#x274C;")), err, config.GetAppConfig().ClientPostgres().Addr())
	}

	client := search.NewClient(config.GetAppConfig().ClientRedisearch().Addr(), "indexer")

	if client == nil {
		return fmt.Errorf("%v initializing redisearch client failed: why? %v\n",
			aurora.Red(html.UnescapeString("&#x274C;")), aurora.Red(err))
	}
	return
}
