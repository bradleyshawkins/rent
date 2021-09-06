package rent

import (
	"fmt"
	"net/http"
)

type Error struct {
	err         error
	code        Code
	message     string
	userMessage string
	fields      []InvalidField
}

type Reason string

const (
	ReasonMissing = "MISSING"
	ReasonInvalid = "INVALID"
)

type InvalidField struct {
	Field  string
	Reason Reason
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
)

const (
	UnauthenticatedMsg = "User is not authenticated"
	InvalidFieldMsg    = "Invalid field was provided"
	InternalMsg        = "Unexpected internal error occurred"
	DuplicateMsg       = "Entity already exists"
	InvalidPayloadMsg  = "Unable to deserialize payload"
	NotExistsMsg       = "Entity does not exist"
)

var codeHttpStatusCodeMap = map[Code]int{
	CodeUnknown:         http.StatusInternalServerError,
	CodeUnauthenticated: http.StatusUnauthorized,
	CodeInvalidField:    http.StatusBadRequest,
	CodeInternal:        http.StatusInternalServerError,
	CodeDuplicate:       http.StatusConflict,
	CodeInvalidPayload:  http.StatusBadRequest,
	CodeNotExists:       http.StatusNotFound,
}

func NewError(err error, options ...ErrorOption) *Error {
	e := &Error{
		err:         err,
		code:        CodeUnknown,
		userMessage: "Unknown error occurred",
	}

	for _, option := range options {
		option(e)
	}

	return e
}

type ErrorOption func(e *Error)

func WithUnauthenticated() ErrorOption {
	return func(e *Error) {
		e.code = CodeUnauthenticated
		e.userMessage = UnauthenticatedMsg
	}
}

func WithInternal() ErrorOption {
	return func(e *Error) {
		e.code = CodeInternal
		e.userMessage = InternalMsg
	}
}

func WithInvalidFields(invalidFields ...InvalidField) ErrorOption {
	return func(e *Error) {
		e.code = CodeInvalidField
		e.userMessage = InvalidFieldMsg
		e.fields = invalidFields
	}
}

func WithDuplicate() ErrorOption {
	return func(e *Error) {
		e.code = CodeDuplicate
		e.userMessage = DuplicateMsg
	}
}

func WithNotExists() ErrorOption {
	return func(e *Error) {
		e.code = CodeNotExists
		e.userMessage = NotExistsMsg
	}
}

func WithInvalidPayload() ErrorOption {
	return func(e *Error) {
		e.code = CodeInvalidPayload
		e.userMessage = InvalidPayloadMsg
	}
}

func WithMessage(message string) ErrorOption {
	return func(e *Error) {
		e.message = message
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: %v, Message: %v", e.err.Error(), e.message)
}

func (e *Error) Code() Code {
	return e.code
}

func (e *Error) UserMessage() string {
	return e.userMessage
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) InvalidFields() []InvalidField {
	return e.fields
}

func (e *Error) HttpStatusCode() int {
	code, ok := codeHttpStatusCodeMap[e.code]
	if !ok {
		return http.StatusInternalServerError
	}
	return code
}
