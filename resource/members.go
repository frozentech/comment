package resource

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/frozentech/api"
	"github.com/frozentech/comment/errors"
	"github.com/frozentech/comment/service"
	"github.com/google/go-github/github"
)

const (
	// MemberURI wallet uri
	MemberURI = "orgs/{organization}/members"
)

// Member initialises Member
type Member struct {
	Resource
	URI string
}

// NewMember initialises Member
func NewMember() Member {
	var (
		resource = Member{
			URI: MemberURI,
			Resource: Resource{
				Name: "comment",
			},
		}
	)

	return resource
}

// AllowedMethods returns the list of allowed methods for this resource
func (me Member) AllowedMethods() []string {
	return []string{http.MethodGet}
}

// Get fetch rule informatiob
func (me Member) Get(ctx context.Context,
	req api.Request,
	w *api.Response,
) (err error) {
	me.Before(req, me.Name)
	defer func() { err = me.After(w, err) }()

	var (
		organization        = req.PathParameters[CommentOrganizationName]
		limit, offset       string
		intLimit, intOffset int
		members             []*github.User
		github              = service.New(ctx, req.Headers[GithubAccessToken])
	)

	if organization == "" {
		err = fmt.Errorf(errors.ErrorResourceNotFound)
		me.Error("Error", errors.NewMessage(errors.ErrorResourceNotFound, CommentOrganizationName))
		return
	}

	if req.QueryStringParameters[QueryLimit] == "" {
		limit = "10"
	}

	if req.QueryStringParameters[QueryOffset] == "" {
		offset = "1"
	}

	if intLimit, err = strconv.Atoi(limit); err != nil {
		err = fmt.Errorf(errors.ErrorInvalidRequestBody)
		me.Error("Error", errors.NewMessage(errors.ErrorInvalidRequestBody, QueryLimit))
		return
	}

	if intOffset, err = strconv.Atoi(offset); err != nil {
		err = fmt.Errorf(errors.ErrorInvalidRequestBody)
		me.Error("Error", errors.NewMessage(errors.ErrorInvalidRequestBody, QueryOffset))
		return
	}

	if members, err = github.GetMember(organization, intLimit, intOffset); err != nil {
		me.Record("Error", err.Error())
		err = fmt.Errorf(errors.ErrorResourceNotFound)
		me.Error("Error", errors.NewMessage(errors.ErrorResourceNotFound, "member"))
		return
	}

	w.Stat(http.StatusOK)
	w.Output(members)

	return nil
}
