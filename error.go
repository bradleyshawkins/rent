package rent

import "fmt"

type Reason string

const (
	Missing Reason = "Missing"
	Invalid Reason = "Invalid"
)

type ValidationError struct {
	Field  string
	Reason Reason
}

func NewValidationError(field string, reason Reason) *ValidationError {
	return &ValidationError{
		Field:  field,
		Reason: reason,
	}
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("Field %s failed validation due to %s", v.Field, v.Reason)
}
