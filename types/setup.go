package types

import "fmt"

type SetupErrorReason string

const (
	SetupNotSet SetupErrorReason = "NotSet"
)

type SetupError struct {
	Field  string
	Reason SetupErrorReason
}

func NewSetupError(field string, reason SetupErrorReason) *SetupError {
	return &SetupError{
		Field:  field,
		Reason: reason,
	}
}

func (v *SetupError) Error() string {
	return fmt.Sprintf("Field %s failed setup validation due to %s", v.Field, v.Reason)
}
