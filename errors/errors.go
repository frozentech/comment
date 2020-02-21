package errors

import (
	"fmt"
	"net/http"

	"github.com/frozentech/array"
)

// custom errors
const (
	ErrorInvalidJSON = "INVALID_JSON_PAYLOAD"

	ErrorInvalidRequestBody = "INVALID_REQUEST_BODY"

	ErrorEmptyRequestBody = "EMPTY_REQUEST_BODY"

	ErrorServiceUnavailable = "SERVICE_UNAVAILABLE"

	ErrorAccessDenied = "ACCESS_DENIED"

	ErrorForbidden = "FORBIDDEN"

	ErrorMethodNotAllowed = "METHOD_NOT_ALLOWED"

	ErrorResourceNotFound = "RESOURCE_NOT_FOUND"

	ErrorInvalidRangeValue = "INVALID_RANGE_VALUE"

	ErrorInvalidTriggerValue = "INVALID_TRIGGER_VALUE"
)

const (
	// StatusOK no error
	StatusOK int = iota

	// StatusInvalidRequest ...
	StatusInvalidRequest

	// StatusAccessDenied ...
	StatusAccessDenied

	// StatusForbidden ...
	StatusForbidden

	// StatusMethodNotAllowed ...
	StatusMethodNotAllowed

	// StatusResourceNotFound ...
	StatusResourceNotFound

	// StatusInvalidJSON ...
	StatusInvalidJSON

	// StatusServiceNotAvailable ...
	StatusServiceNotAvailable
)

// ErrorMessages ...
var ErrorMessages = map[string]ErrDetail{
	ErrorResourceNotFound: ErrDetail{
		Code:    StatusResourceNotFound,
		Message: http.StatusText(http.StatusNotFound),
		Hint:    "Resource not found",
		Status:  http.StatusNotFound,
	},
	ErrorForbidden: ErrDetail{
		Code:    StatusForbidden,
		Message: http.StatusText(http.StatusForbidden),
		Hint:    "Unauthorized",
		Status:  http.StatusForbidden,
	},
	ErrorAccessDenied: ErrDetail{
		Code:    StatusAccessDenied,
		Message: http.StatusText(http.StatusForbidden),
		Hint:    "Unauthorized",
		Status:  http.StatusForbidden,
	},
	ErrorMethodNotAllowed: ErrDetail{
		Code:    StatusMethodNotAllowed,
		Message: http.StatusText(http.StatusMethodNotAllowed),
		Hint:    "Method Not Allowed",
		Status:  http.StatusMethodNotAllowed,
	},
	ErrorInvalidJSON: ErrDetail{
		Code:    StatusInvalidJSON,
		Message: http.StatusText(http.StatusUnprocessableEntity),
		Hint:    "Request body is not a valid json",
		Status:  http.StatusUnprocessableEntity,
	},
	ErrorInvalidRequestBody: ErrDetail{
		Code:    StatusInvalidRequest,
		Message: http.StatusText(http.StatusBadRequest),
		Hint:    "Request body contains invalid value",
		Status:  http.StatusBadRequest,
	},
	ErrorEmptyRequestBody: ErrDetail{
		Code:    StatusInvalidRequest,
		Message: http.StatusText(http.StatusBadRequest),
		Hint:    "Request body contains invalid value",
		Status:  http.StatusBadRequest,
	},
	ErrorInvalidRangeValue: ErrDetail{
		Code:    StatusInvalidRequest,
		Message: http.StatusText(http.StatusBadRequest),
		Hint:    "Request body contains invalid value",
		Status:  http.StatusBadRequest,
	},
	ErrorServiceUnavailable: ErrDetail{
		Code:    StatusServiceNotAvailable,
		Message: http.StatusText(http.StatusServiceUnavailable),
		Hint:    "A unexpected error occured",
		Status:  http.StatusServiceUnavailable,
	},
}

// ErrDetail ...
type ErrDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Hint    string `json:"hint"`
	Status  int    `json:"-"`
}

// ListErrors ...
func ListErrors() []ErrDetail {
	var list []ErrDetail
	var listError = make(map[string]string)

	for _, err := range ErrorMessages {
		if !array.KeyExist(fmt.Sprint(err.Code), listError) {
			list = append(list, err)
			listError[fmt.Sprint(err.Code)] = err.Message
		}
	}

	return list

}

// New ErrDetail
func New(code string, hint string) ErrDetail {
	var (
		err ErrDetail
		ok  bool
	)

	if err, ok = ErrorMessages[code]; !ok {
		err = ErrorMessages[ErrorServiceUnavailable]
	}

	if hint != "" {
		err.Hint = hint
	}

	return err
}
