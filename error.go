package rent

import (
	"fmt"
)

type ValidationReason string

const (
	Missing ValidationReason = "Missing"
	Invalid ValidationReason = "Invalid"
)

type ValidationError struct {
	Field  string
	Reason ValidationReason
}

func NewValidationError(field string, reason ValidationReason) *ValidationError {
	return &ValidationError{
		Field:  field,
		Reason: reason,
	}
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("Field %s failed validation due to %s", v.Field, v.Reason)
}
