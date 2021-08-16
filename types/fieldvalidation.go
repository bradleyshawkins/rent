package types

import "fmt"

type FieldValidationReason string

const (
	Missing FieldValidationReason = "Missing"
	Invalid FieldValidationReason = "Invalid"
)

type FieldValidationError struct {
	Field  string
	Reason FieldValidationReason
}

func NewFieldValidationError(field string, reason FieldValidationReason) *FieldValidationError {
	return &FieldValidationError{
		Field:  field,
		Reason: reason,
	}
}

func (v *FieldValidationError) Error() string {
	return fmt.Sprintf("Field %s failed validation due to %s", v.Field, v.Reason)
}
