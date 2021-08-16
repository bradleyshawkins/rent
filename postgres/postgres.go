package postgres

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

type Postgres struct {
	db *sqlx.DB
}

func New(connectionString string, migrationPath string) (*Postgres, error) {
	log.Println("Connecting to database...")
	db, err := sqlx.Open("pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to open database connection. Error: %v", err)
	}

	p := &Postgres{db: db}
	if err := p.Migrate(migrationPath); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Postgres) Migrate(migrationPath string) error {
	log.Println("Running database migrations...")
	goose.SetTableName("goose_db_version")
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("unable to set goose dialect to postgres. Error: %v", err)
	}
	err = goose.Up(p.db.DB, migrationPath)
	if err != nil {
		return fmt.Errorf("unable to migrate database. Error: %v", err)
	}
	return nil
}
