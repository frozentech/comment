package service

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Service ...
type Service struct {
	Context context.Context
	Token   string
	Client  *github.Client
}

// New ...
func New(ctx context.Context, accessToken string) *Service {

	s := Service{
		Context: ctx,
		Token:   accessToken,
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: s.Token})
	tc := oauth2.NewClient(ctx, ts)

	s.Client = github.NewClient(tc)

	return &s
}

// GetOrganization ....
func (me *Service) GetOrganization(name string) (organization *github.Organization, err error) {
	organization, _, err = me.Client.Organizations.Get(me.Context, name)
	if err != nil {
		return
	}

	return
}

// GetMember ...
func (me *Service) GetMember(name string, limit int, offset int) (members []*github.User, err error) {
	option := github.ListMembersOptions{}
	option.Page = offset
	option.PerPage = limit
	members, _, err = me.Client.Organizations.ListMembers(me.Context, name, &option)
	if err != nil {
		return
	}

	return
}

// GetLoggedInUser ...
func (me *Service) GetLoggedInUser() (user *github.User, err error) {
	user, _, err = me.Client.Users.Get(me.Context, "")
	return
}

// GetCheckMembership ...
func (me *Service) GetCheckMembership(name string, user string) bool {
	isMember, _, _ := me.Client.Organizations.IsMember(me.Context, name, user)
	return isMember
}
