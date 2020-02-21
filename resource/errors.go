package resource

import (
	"context"
	"net/http"

	"github.com/frozentech/api"
	commentError "github.com/frozentech/comment/errors"
)

const (
	// ErrorsURI wallet uri
	ErrorsURI = "/error"
)

// Errors initialises Errors
type Errors struct {
	Resource
	URI string
}

// NewErrors initialises Errors
func NewErrors() Errors {
	var (
		resource = Errors{
			URI: ErrorsURI,
			Resource: Resource{
				Name: "error",
			},
		}
	)

	return resource
}

// AllowedMethods returns the list of allowed methods for this resource
func (me Errors) AllowedMethods() []string {
	return []string{http.MethodGet}
}

// Get fetch error informatiob
func (me Errors) Get(ctx context.Context,
	req api.Request,
	w *api.Response,
) (err error) {
	me.Before(req, me.Name)
	defer func() { err = me.After(w, err) }()

	w.Stat(http.StatusOK)
	w.Output(struct {
		Count  int                      `json:"count"`
		Errors []commentError.ErrDetail `json:"errors"`
	}{
		Count:  len(commentError.ErrorMessages),
		Errors: commentError.ListErrors(),
	})

	return nil
}
