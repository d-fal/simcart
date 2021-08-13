package postgres

import (
	"simcart/config"
	"sync"

	"github.com/go-pg/pg/v10"
)

var (
	Storage store = &psql{}
	once    sync.Once
)

type Tx interface {
	Begin() error
	Commit() error
	Rollback() error
	Close() error
	Get() *pg.Tx
	Db() *pg.DB
}

// store interface is interface for store things into postgres
type store interface {
	Connect(level config.LogLevel, config *config.Postgres) error
	Get() (*pg.DB, error)
	Transaction() (Tx, error)
	Close()
}

// postgres struct
type psql struct {
	db *pg.DB
}

type tx struct {
	db  *pg.DB
	tx  *pg.Tx
	err error
}
type Hooks struct {
	afterSelect int
	verbose     config.LogLevel
}
