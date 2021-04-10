package mysql

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

type MySQL struct {
	db *sqlx.DB
}

func New(connectionString string, migrationPath string) (*MySQL, error) {
	log.Println("Connecting to database...")
	db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to open database connection. Error: %v", err)
	}

	log.Println("Running database migrations...")
	goose.SetTableName("goose_db_version")
	err = goose.SetDialect("mysql")
	if err != nil {
		return nil, fmt.Errorf("unable to set goose dialect to sql. Error: %v", err)
	}
	err = goose.Up(db.DB, migrationPath)
	if err != nil {
		return nil, fmt.Errorf("unable to migrate database. Error: %v", err)
	}

	return &MySQL{db: db}, nil
}
