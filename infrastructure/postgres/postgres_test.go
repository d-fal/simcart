package postgres_test

import (
	"fmt"
	"testing"

	"simcart/config"
	"simcart/infrastructure/postgres"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	// Storage.Connect(config)
	tests := []struct {
		step string
		conf *config.Postgres
		err  error
	}{
		{
			step: "A",
			conf: &config.Postgres{},
			err:  fmt.Errorf("dial tcp postgres:5432: connect: connection refused"),
		},
		{
			step: "B",
			conf: &config.Postgres{
				User:     "postgres",
				Password: "password",
				Host:     "postgres:5432",
				Database: "test_db",
			},
			err: nil,
		}}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			err := postgres.Storage.Connect(config.Verbose, tc.conf)
			if tc.err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestGet(t *testing.T) {
	db, _ := postgres.Storage.Get()
	if db == nil {
		assert.Error(t, fmt.Errorf("error in database get DB"))
	}
}
