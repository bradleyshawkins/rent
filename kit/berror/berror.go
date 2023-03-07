package berror

import (
	"fmt"
	"net/http"
)

type Error struct {
	err     error
	code    Code
	message string
}

type Code int

const (
	CodeUnknown Code = iota + 1
	CodeUnauthenticated
	CodeInvalidPayload
	CodeInvalidField
	CodeInternal
	CodeDuplicate
	CodeNotExists
	CodeRequiredEntityNotExists
	CodeDisabled
)

var codeHttpStatusCodeMap = map[Code]int{
	CodeUnknown:                 http.StatusInternalServerError,
	CodeUnauthenticated:         http.StatusUnauthorized,
	CodeInvalidField:            http.StatusBadRequest,
	CodeInternal:                http.StatusInternalServerError,
	CodeDuplicate:               http.StatusConflict,
	CodeInvalidPayload:          http.StatusBadRequest,
	CodeNotExists:               http.StatusNotFound,
	CodeRequiredEntityNotExists: http.StatusConflict,
	CodeDisabled:                http.StatusNotFound,
}

func Internal(err error, msg string) *Error {
	return NewError(err, CodeInternal, msg)
}

func Duplicate(err error, msg string) *Error {
	return NewError(err, CodeDuplicate, msg)
}

func NotFound(err error, msg string) *Error {
	return NewError(err, CodeNotExists, msg)
}

func InvalidPayload(err error, msg string) *Error {
	return NewError(err, CodeInternal, msg)
}

func InvalidField(err error, msg string) *Error {
	return NewError(err, CodeInvalidField, msg)
}

func NewError(err error, code Code, msg string) *Error {
	return &Error{
		err:     err,
		code:    code,
		message: msg,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: %v, Message: %v", e.err.Error(), e.message)
}

func (e *Error) Code() Code {
	return e.code
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) HttpStatusCode() int {
	code, ok := codeHttpStatusCodeMap[e.code]
	if !ok {
		return http.StatusInternalServerError
	}
	return code
}

func (e *Error) WriteResponse(w http.ResponseWriter) {
	http.Error(w, e.Message(), e.HttpStatusCode())
}
