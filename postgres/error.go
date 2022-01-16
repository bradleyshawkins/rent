package postgres

import (
	"database/sql"
	"errors"
	"fmt"
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
			return rent.NewError(err, rent.WithDuplicate(), rent.WithMessage(fmt.Sprintf("duplicate entry found. Details: %s", pgErr.Detail)))
		case foreignKeyFailed:
			return rent.NewError(err, rent.WithRequiredEntityNotExist(), rent.WithMessage(fmt.Sprintf("required entity does not exist. Details: %s", pgErr.Detail)))
		}

	}
	return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unexpected error occurred"))
}
