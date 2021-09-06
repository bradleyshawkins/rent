package postgres

import (
	"database/sql"
	"errors"

	"github.com/jackc/pgconn"

	"github.com/bradleyshawkins/rent"
)

const (
	duplicateEntry string = "23505"
)

func toRentError(err error) error {
	if err == sql.ErrNoRows {
		return rent.NewError(err, rent.WithNotExists(), rent.WithMessage("entity does not exist"))
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case duplicateEntry:
			return rent.NewError(err, rent.WithDuplicate(), rent.WithMessage("duplicate entry found"))
		}
	}
	return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unexpected error occurred"))
}
