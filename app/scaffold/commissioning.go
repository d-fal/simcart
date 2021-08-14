package scaffold

import (
	"context"
	"fmt"
	"html"
	"simcart/config"
	"simcart/infrastructure/postgres"
	"simcart/infrastructure/redis"
	"simcart/infrastructure/search"

	"github.com/logrusorgru/aurora"
)

func (s *skeleton) commissioning(ctx context.Context) (err error) {

	s.context, s.cancel = context.WithCancel(ctx)

	s.tracer, s.closer, err = s.tracing()

	if err != nil {
		return err
	}

	logLevel := config.Silent

	if s.params.Mode() {
		logLevel = config.Verbose
	}

	if err = postgres.Storage.Connect(logLevel,
		config.GetAppConfig().ClientPostgres()); err != nil {
		return fmt.Errorf("%v \t %v\n", aurora.White(html.UnescapeString("&#x274C;")), err)
	}

	if err := s.broker(); err != nil {
		return err
	}

	if err := redis.NewClient( /* 0 th db is selecte*/ 0).Connect(config.GetAppConfig()); err != nil {

		fmt.Printf("%v starting redis failed: why? %v\n",
			aurora.Red(html.UnescapeString("&#x274C;")), aurora.Red(err))

		return err
	}

	client := search.NewClient(config.GetAppConfig().ClientRedisearch().Addr(), "indexer")

	if client == nil {
		return fmt.Errorf("%v initializing redisearch client failed: why? %v\n",
			aurora.Red(html.UnescapeString("&#x274C;")), aurora.Red(err))
	}
	return
}
