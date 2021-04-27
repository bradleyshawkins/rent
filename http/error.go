package http

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Error struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

func NewError(statusCode int, message string) *Error {
	return &Error{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (e *Error) WriteError(w http.ResponseWriter) {
	w.WriteHeader(e.StatusCode)
	_ = json.NewEncoder(w).Encode(e)
}

func isDuplicate(err error) bool {
	type isDuplicate interface {
		IsDuplicate() bool
	}

	e := errors.Unwrap(err)
	var i isDuplicate
	return errors.As(e, &i) && i.IsDuplicate()
}
