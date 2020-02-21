package resource

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/frozentech/api"
	"github.com/frozentech/comment/database"
	"github.com/frozentech/comment/errors"
	"github.com/frozentech/comment/model"
	"github.com/frozentech/comment/service"
)

const (
	// CommentURI wallet uri
	CommentURI = "orgs/{organization}/comments"

	// CommentOrganizationName ...
	CommentOrganizationName = "organization"
)

// Comment initialises Comment
type Comment struct {
	Resource
	URI string
}

// NewComment initialises Comment
func NewComment() Comment {
	var (
		resource = Comment{
			URI: CommentURI,
			Resource: Resource{
				Name: "comment",
			},
		}
	)

	return resource
}

// AllowedMethods returns the list of allowed methods for this resource
func (me Comment) AllowedMethods() []string {
	return []string{http.MethodGet, http.MethodPost, http.MethodDelete}
}

// Get comments all comments
func (me Comment) Get(ctx context.Context,
	req api.Request,
	w *api.Response,
) (err error) {
	me.Before(req, me.Name)
	defer func() { err = me.After(w, err) }()

	var (
		organization = req.PathParameters[CommentOrganizationName]
		comments     []model.Comment
		github       = service.New(ctx, req.Headers[GithubAccessToken])
	)

	if organization == "" {
		err = fmt.Errorf(errors.ErrorResourceNotFound)
		me.Error("Error", errors.NewMessage(errors.ErrorResourceNotFound, CommentOrganizationName))
		return
	}

	_, err = github.GetOrganization(organization)
	if err != nil {
		me.Record("Error", err.Error())
		err = fmt.Errorf(errors.ErrorResourceNotFound)
		me.Error("Error", errors.NewMessage(errors.ErrorResourceNotFound, CommentOrganizationName))
		return
	}

	if comments, err = database.GetComment(ctx, organization); err != nil {
		me.Record("Error", err.Error())
		err = fmt.Errorf(errors.ErrorResourceNotFound)
		me.Error("Error", errors.NewMessage(errors.ErrorResourceNotFound, "comment"))
		return
	}

	w.Stat(http.StatusOK)
	w.Output(comments)

	return nil
}

// Post comments to a organization
func (me Comment) Post(ctx context.Context,
	req api.Request,
	w *api.Response,
) (err error) {
	me.Before(req, me.Name)
	defer func() { err = me.After(w, err) }()

	var (
		payload      model.Request
		organization = req.PathParameters[CommentOrganizationName]
		comment      *model.Comment
		github       = service.New(ctx, req.Headers[GithubAccessToken])
	)

	if req.Headers[GithubAccessToken] == "" {
		err = fmt.Errorf(errors.ErrorAccessDenied)
		me.Error("Error", errors.NewMessage(errors.ErrorAccessDenied, ""))
		return
	}

	if organization == "" {
		err = fmt.Errorf(errors.ErrorResourceNotFound)
		me.Error("Error", errors.NewMessage(errors.ErrorResourceNotFound, CommentOrganizationName))
		return
	}

	_, err = github.GetOrganization(organization)
	if err != nil {
		me.Record("Error", err.Error())
		err = fmt.Errorf(errors.ErrorResourceNotFound)
		me.Error("Error", errors.NewMessage(errors.ErrorResourceNotFound, CommentOrganizationName))
		return
	}

	if err = json.Unmarshal([]byte(req.Body), &payload); err != nil {
		me.Record("Error", err.Error())
		err = fmt.Errorf(errors.ErrorInvalidRequestBody)
		me.Error("Error", errors.NewMessage(errors.ErrorInvalidRequestBody, "body"))
		return
	}

	member, err := github.GetLoggedInUser()
	if err != nil {
		me.Record("Error", err.Error())
		err = fmt.Errorf(errors.ErrorAccessDenied)
		me.Error("Error", errors.NewMessage(errors.MessageAccessDenied, ""))
		return
	}

	payload.Member = model.Member{
		ID:      member.GetID(),
		Name:    member.GetLogin(),
		Profile: member.GetAvatarURL(),
	}

	if os.Getenv(MembershipFilter) != "" {
		if !github.GetCheckMembership(organization, payload.Member.Name) {
			err = fmt.Errorf(errors.ErrorAccessDenied)
			me.Error("Error", errors.NewMessage(errors.MessageAccessDenied, ""))
			return
		}
	}

	comment = model.NewComment()
	comment.Member = payload.Member
	comment.Message = payload.Message

	if err = database.PostComment(ctx, *comment); err != nil {
		me.Record("Error", err.Error())
		err = fmt.Errorf(errors.ErrorResourceNotFound)
		me.Error("Error", errors.NewMessage(errors.ErrorResourceNotFound, "comment"))
		return
	}

	w.Stat(http.StatusCreated)
	w.Output(comment)

	return nil
}

// Delete all comments from a organization
func (me Comment) Delete(ctx context.Context,
	req api.Request,
	w *api.Response,
) (err error) {
	me.Before(req, me.Name)
	defer func() { err = me.After(w, err) }()

	var (
		organization = req.PathParameters[CommentOrganizationName]
		github       = service.New(ctx, req.Headers[GithubAccessToken])
	)

	if req.Headers[GithubAccessToken] == "" {
		err = fmt.Errorf(errors.ErrorAccessDenied)
		me.Error("Error", errors.NewMessage(errors.ErrorAccessDenied, ""))
		return
	}

	if organization == "" {
		err = fmt.Errorf(errors.ErrorResourceNotFound)
		me.Error("Error", errors.NewMessage(errors.ErrorResourceNotFound, CommentOrganizationName))
		return
	}

	_, err = github.GetOrganization(organization)
	if err != nil {
		me.Record("Error", err.Error())
		err = fmt.Errorf(errors.ErrorResourceNotFound)
		me.Error("Error", errors.NewMessage(errors.ErrorResourceNotFound, CommentOrganizationName))
		return
	}

	if err = database.DeleteComment(ctx, organization); err != nil {
		me.Record("Error", err.Error())
		err = fmt.Errorf(errors.ErrorResourceNotFound)
		me.Error("Error", errors.NewMessage(errors.ErrorResourceNotFound, CommentOrganizationName))
		return
	}

	w.Stat(http.StatusNoContent)
	w.Output(nil)

	return nil
}
