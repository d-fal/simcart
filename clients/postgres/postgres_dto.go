package postgres

import (
	"context"
	"fmt"
	"net"
	"simcart/config"
	"simcart/pkg/logger"

	"github.com/go-pg/pg/v10"
	"github.com/logrusorgru/aurora"
	"go.uber.org/zap"
)

// Connect method job is connect to postgres database and check migration
func (p *psql) Connect(level config.LogLevel, cnf *config.Postgres) error {
	var err error

	once.Do(func() {
		p.db = pg.Connect(&pg.Options{
			User:     cnf.User,
			Password: cnf.Password,
			Database: cnf.Database,
			Addr: net.JoinHostPort(
				cnf.Host,
				cnf.Port,
			),
		})
		if err = p.db.Ping(context.Background()); err != nil {
			lg := logger.
				NewPrototype()

			lg.Development().
				Level(zap.ErrorLevel).Commit("init configs")
			return
		}

		hooks := new(Hooks)

		hooks.verbose = level

		p.db.AddQueryHook(hooks)

	})

	return err
}

func (p *psql) Get() (*pg.DB, error) {
	if p.db == nil {
		return nil, fmt.Errorf("db is not initialized")
	}
	return p.db, nil
}

func (p *psql) Close() {
	if p != nil {
		p.Close()
	}
}

func (p *psql) Store() {}

func (p *psql) Transaction() (Tx, error) {

	t := new(tx)

	pgdb, err := p.Get()

	if err != nil {
		return t, err
	}

	t.db = pgdb

	return t, nil
}

func (t *tx) Begin() error {
	tx, err := t.db.Begin()
	t.tx = tx
	return err
}

func (t *tx) Commit() error {
	return t.tx.Commit()
}

func (t *tx) Close() error {
	return t.tx.Close()
}

func (t *tx) Rollback() error {
	return t.tx.Rollback()
}

func (t *tx) Get() *pg.Tx {
	return t.tx
}

func (t *Hooks) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (t *Hooks) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	switch t.verbose {
	case config.Verbose:
		query, _ := q.FormattedQuery()
		fmt.Printf("%v\t%v\n", aurora.White("go-pg").BgGreen(), aurora.Cyan(string(query)))
	}

	return nil
}

func (t *tx) Db() *pg.DB {
	return t.db
}
