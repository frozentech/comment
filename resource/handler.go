package resource

import "github.com/frozentech/api"

// NewHandler ...
func NewHandler() api.Handlers {

	commentResource := NewComment()
	memberResource := NewMember()

	return api.Handlers{
		commentResource.URI: commentResource,
		memberResource.URI:  memberResource,
	}
}
