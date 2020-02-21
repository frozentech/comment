package resource

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/frozentech/api"
	commentError "github.com/frozentech/comment/errors"
	"github.com/frozentech/logs"
)

const (

	// TestingEnv TESTING
	TestingEnv = "TESTING"

	// GithubAccessToken ACCESS TOKEN
	GithubAccessToken = "x-access-token"

	// MembershipFilter MEMBERSHIP_FILTER
	MembershipFilter = "MEMBERSHIP_FILTER"

	// QueryLimit limit
	QueryLimit = "limit"

	// QueryOffset offset
	QueryOffset = "offset"
)

// Resource struct
type Resource struct {
	api.Resource
	Name string
	Hint string
}

// After after module execution
func (me *Resource) After(w *api.Response, err error) error {

	_, fn, line, _ := runtime.Caller(2)

	rec := recover()
	if rec != nil {
		me.Record("Recovery", rec)
	}

	if err != nil {

		me.Record("Line", line)
		me.Record("File", fn[strings.LastIndex(fn, "/")+1:])

		if me.Hint != "" {
			me.Record("Hint", me.Hint)
		}

		e := commentError.New(err.Error(), me.Hint)

		w.Stat(e.Status)
		w.Output(e)

	}

	me.Record("Response", w.Body)
	me.Record("Status", w.StatusCode)

	return nil
}

// Record return metrics
func (me *Resource) Record(header string, message interface{}) {
	logs.Story.Record(header, message)
}

func (me *Resource) Error(header string, message interface{}) {
	me.Hint = fmt.Sprint(message)
	logs.Story.Record(header, message)
}

// Before logs the api request
func (me *Resource) Before(req api.Request, api string) {
	me.Record("API", api)
	me.Record("URL", req.Path)
	me.Record("Method", req.HTTPMethod)
	me.Record("Request", req.Body)
	me.Record("Headers", req.Headers)
	me.Record("Key", req.GetResourceKey())
	me.Record("URI", req.PathParameters)
	me.Record("Payload", req.QueryStringParameters)
}
