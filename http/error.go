package http

import (
	"encoding/json"
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
