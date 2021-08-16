package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
)

const (
	duplicateEntry string = "23505"
)

type Error struct {
	err       error
	message   string
	errorType errorType
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: %s, Message: %s", e.err, e.message)
}

func (e *Error) IsDuplicate() bool {
	return e.errorType == duplicate
}

func (e *Error) IsNotExists() bool {
	return e.errorType == notExists
}

func convertToError(err error, message string) *Error {

	if err == sql.ErrNoRows {
		return &Error{
			err:       err,
			message:   message,
			errorType: notExists,
		}
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case duplicateEntry:
			return &Error{
				err:       err,
				message:   message,
				errorType: duplicate,
			}
		}
	}

	return &Error{
		err:       err,
		message:   message,
		errorType: unknown,
	}
}

type errorType int

const (
	unknown errorType = iota
	duplicate
	notExists
)
