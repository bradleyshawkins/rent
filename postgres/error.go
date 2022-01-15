package postgres

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"

	"github.com/bradleyshawkins/rent"
)

const (
	foreignKeyFailed string = "23503"
	duplicateEntry   string = "23505"
)

func toRentError(err error) error {
	if err == sql.ErrNoRows {
		return rent.NewError(err, rent.WithNotExists(), rent.WithMessage("entity does not exist"))
	}

	var pgErr *pq.Error
	if errors.As(err, &pgErr) {
		switch string(pgErr.Code) {
		case duplicateEntry:
			return rent.NewError(err, rent.WithDuplicate(), rent.WithMessage("duplicate entry found"))
		case foreignKeyFailed:
			return rent.NewError(err, rent.WithRequiredEntityNotExist(), rent.WithMessage("required entity does not exist"))
		}

	}
	return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unexpected error occurred"))
}
