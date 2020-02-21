package errors

import (
	"fmt"
	"strings"
)

// Messages ...
const (
	MessageNotFound     = "Record not found"
	MessageInvalid      = "Parameter is invalid"
	MessageRange        = "Parameter must be a valid range value"
	MessageAccessDenied = "Access denied"
	MessageRequired     = "Parameter is a required field"
)

// NewMessage return formatted error message
func NewMessage(err string, field string) error {
	switch err {
	case ErrorResourceNotFound:
		return fmt.Errorf("%s : %s", strings.ToLower(field), MessageNotFound)
	case ErrorInvalidRequestBody:
		return fmt.Errorf("%s : %s", strings.ToLower(field), MessageInvalid)
	case ErrorInvalidRangeValue:
		return fmt.Errorf("%s : %s", strings.ToLower(field), MessageRange)
	case ErrorEmptyRequestBody:
		return fmt.Errorf("%s : %s", strings.ToLower(field), MessageRequired)
	case ErrorAccessDenied:
		return fmt.Errorf("%s", MessageAccessDenied)
	}
	return fmt.Errorf("%s : %s", strings.ToLower(field), err)
}
