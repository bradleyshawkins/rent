package postgres

import (
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
)

const (
	duplicateEntry string = "23505"
)

func convertToError(err error, message string) *Error {

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case duplicateEntry:
			return &Error{
				err:         err,
				message:     message,
				isDuplicate: true,
			}
		}
	}

	return &Error{
		err:         err,
		message:     message,
		isDuplicate: false,
	}
}

type Error struct {
	err         error
	message     string
	isDuplicate bool
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: %s, Message: %s", e.err, e.message)
}

func (e *Error) IsDuplicate() bool {
	return e.isDuplicate
}
