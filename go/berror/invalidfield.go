package berror

import "errors"

type InvalidFieldError struct {
	*Error
	fields []InvalidField
}

type Reason string

const (
	ReasonMissing = "MISSING"
	ReasonInvalid = "INVALID"
)

type InvalidField struct {
	Field  string
	Reason Reason
	Value  string
}

func NewInvalidFieldsError(msg string, invalidFields ...InvalidField) *InvalidFieldError {
	return &InvalidFieldError{
		Error: &Error{
			err:         errors.New(msg),
			code:        CodeInvalidField,
			message:     "",
			userMessage: "",
		},
		fields: invalidFields,
	}
}

func (i *InvalidFieldError) InvalidFields() []InvalidField {
	return i.fields
}
