package postgres_test

// import (
// 	"fmt"
// 	"simcart/config"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestConnect(t *testing.T) {
// 	// Storage.Connect(config)
// 	tests := []struct {
// 		step string
// 		conf config.Config
// 		err  error
// 	}{
// 		{
// 			step: "A",
// 			conf: config.Config{
// 				Debug: true,
// 			},
// 			err: fmt.Errorf("dial tcp 127.0.0.1:5432: connect: connection refused"),
// 		},
// 		{
// 			step: "B",
// 			conf: config.Config{
// 				Debug: true,
// 				POSTGRES: config.Database{
// 					Username: "admin",
// 					Password: "password",
// 					Host:     "127.0.0.1:5432",
// 					Schema:   "schema",
// 				},
// 			},
// 			err: nil,
// 		}}

// 	for _, tc := range tests {
// 		t.Run(tc.step, func(t *testing.T) {
// 			err := Storage.Connect(tc.conf)
// 			if tc.err != nil {
// 				assert.Equal(t, tc.err.Error(), err.Error())
// 			}
// 		})
// 	}
// }

// func TestGet(t *testing.T) {
// 	db, _ := Storage.Get()
// 	if db == nil {
// 		assert.Error(t, fmt.Errorf("error in database get DB"))
// 	}
// }
