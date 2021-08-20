package scaffold

import (
	"context"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {

	scf := NewAppScaffold()

	scf.SetConfigPath("../../")

	setup := scf.Start(context.TODO(), false, WithRedisearch(), WithPostgres())

	t.Run("setup", func(t *testing.T) {
		if err := setup(); err != nil {
			assert.Error(t, err)
			return
		}

	})

	t.Run("options", func(t *testing.T) {
		jg := WithOpenTracing()
		jg(context.TODO(), scf.(*skeleton))

		rds := WithRedisearch()
		rds(context.TODO(), scf.(*skeleton))

		pg := WithPostgres()
		pg(context.TODO(), scf.(*skeleton))

	})

	t.Run("check server", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		t.Run("start server", func(t *testing.T) {
			go scf.Start(ctx, true)
		})

		t.Run("stop server", func(t *testing.T) {
			time.Sleep(time.Second * 2)
			cancel()

		})
	})
}
