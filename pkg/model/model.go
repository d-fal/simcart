package model

import (
	"time"

	"github.com/go-pg/pg/v10"
)

type Model struct {
	Id uint64
}

type Deleteables struct {
	CreatedAt time.Time `pg:"default:now()"`
	DeletedAt time.Time `pg:",soft_delete"`
}

type InsertFunc func(tx *pg.Tx) error
type SelectOrInsertFunc func(db *pg.DB) error
type SelectFunc func(db *pg.DB) error
