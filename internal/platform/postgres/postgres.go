package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

var (
	ErrNotFound       = sql.ErrNoRows
	ErrDuplicateEntry = errors.New("duplicate entry")
)

// https://github.com/lib/pq/blob/master/error.go
const (
	uniqueViolation = "23505"
)

type Database struct {
	db   *sqlx.DB
	conn sqlx.ExtContext
	inTx bool
}

type Config struct {
	Host     string
	Database string
	Schema   string
	Username string
	Password string
}

func (c *Config) String() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable", c.Username, c.Password, c.Host, c.Database)
}

func New(c *Config) (*Database, error) {
	db, err := sqlx.Open("postgres", c.String())
	if err != nil {
		return nil, err
	}

	return &Database{
		db:   db,
		conn: db,
	}, nil
}

func (d *Database) InTx(ctx context.Context, f func(tx *Database) error) error {
	if d.inTx {
		return f(d)
	}

	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx %w", err)
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			if errors.Is(err, sql.ErrTxDone) {
				return
			}
		}
	}()

	err = f(&Database{
		db:   d.db,
		conn: tx,
		inTx: true,
	})
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit tx %w", err)
	}

	return nil
}

func (d *Database) NamedExecContext(ctx context.Context, query string, args ...any) error {
	_, err := sqlx.NamedExecContext(ctx, d.conn, query, args)
	if err != nil {
		return classifyError(err)
	}
	return nil
}

func (d *Database) GetContext(ctx context.Context, dest any, query string, args ...any) error {
	err := sqlx.GetContext(ctx, d.conn, dest, query, args)
	if err != nil {
		return classifyError(err)
	}

	return nil
}

func (d *Database) SelectContext(ctx context.Context, dest any, query string, args ...any) error {
	err := sqlx.SelectContext(ctx, d.conn, dest, query, args)
	if err != nil {
		return classifyError(err)
	}
	return nil
}

func classifyError(err error) error {
	if err == sql.ErrNoRows {
		return ErrNotFound
	}
	if pErr, ok := err.(pq.Error); ok {
		if pErr.Code == uniqueViolation {
			return ErrDuplicateEntry
		}
	}
	return err
}
