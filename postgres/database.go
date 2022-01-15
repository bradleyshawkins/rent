package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(connectionString, migrationPath string) (*Database, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	database := &Database{db: db}
	err = database.migrate(migrationPath)
	if err != nil {
		return nil, err
	}
	return database, nil
}

func (d *Database) migrate(migrationPath string) error {
	log.Println("Running database migrations...")
	goose.SetTableName("goose_db_version")
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("unable to set goose dialect to postgres. Error: %v", err)
	}
	err = goose.Up(d.db, migrationPath)
	if err != nil {
		return fmt.Errorf("unable to migrate database. Error: %v", err)
	}
	return nil
}

func (d *Database) begin() (*transaction, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return nil, err
	}

	return &transaction{tx: tx}, nil
}

type transaction struct {
	tx *sql.Tx
}

func (t *transaction) commit() error {
	err := t.tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (t *transaction) rollback() error {
	return t.tx.Rollback()
}
