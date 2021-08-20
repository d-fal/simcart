package config_test

import (
	"simcart/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {

	cnf := config.GetAppConfig()

	cnf.SetPath("..")
	cnf.Debug(true)
	cnf.Level(uint8(config.Silent))

	t.Run("redisearch", func(t *testing.T) {
		rds := cnf.ClientRedisearch()

		assert.NotNil(t, rds.Addr())
	})

	t.Run("jaeger", func(t *testing.T) {
		jg := cnf.ClientJaeger()

		assert.NotNil(t, jg.Addr())
	})

	t.Run("postgres", func(t *testing.T) {
		pg := cnf.ClientPostgres()

		assert.NotNil(t, pg.Addr())
	})

	t.Run("server", func(t *testing.T) {
		gp := cnf.ServerGRPC()

		assert.NotNil(t, gp.Addr())

		rst := cnf.ServerRest()

		assert.NotNil(t, rst.Addr())
	})

	t.Run("debug", func(t *testing.T) {
		assert.Equal(t, cnf.Mode(), true)

		cnf.Debug(false)

		assert.Equal(t, cnf.Mode(), false)
	})
	t.Run("application", func(t *testing.T) {
		assert.NotNil(t, cnf.Application())
	})
}
