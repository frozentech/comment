package resource_test

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/frozentech/api"
	"github.com/frozentech/comment/resource"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func NewTestMember() resource.Member {
	return resource.NewMember()
}

func NewMemberRequest(payload string, method string, path map[string]string, query map[string]string) events.APIGatewayProxyRequest {
	member := NewTestMember()

	return events.APIGatewayProxyRequest{
		Resource: member.URI,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		PathParameters:        path,
		QueryStringParameters: query,
		HTTPMethod:            method,
		Body:                  payload,
	}
}

func TestMember_New(t *testing.T) {
	a := NewTestMember()
	assert := assert.New(t)
	assert.IsType(resource.Member{}, a)
	assert.Equal("orgs/{organization}/members", a.URI)
}

func TestMember_Methods(t *testing.T) {
	a := NewTestMember()
	assert := assert.New(t)
	assert.Equal(a.AllowedMethods(), []string{http.MethodGet})
}

func TestMember_Get_Not_Found(t *testing.T) {
	handler := resource.NewHandler()
	request := NewMemberRequest("", http.MethodGet, map[string]string{resource.CommentOrganizationName: ""}, map[string]string{})
	res, err := api.NewHandler(handler)(GetContext(), request)

	require := require.New(t)
	require.Nil(err)

	require.EqualValues(http.StatusNotFound, res.StatusCode)
}

func TestMember_Get_Bad_Request_Organization_Error(t *testing.T) {
	handler := resource.NewHandler()
	request := NewMemberRequest("", http.MethodGet, map[string]string{resource.CommentOrganizationName: "xendit"}, map[string]string{})
	res, err := api.NewHandler(handler)(GetContext(), request)

	require := require.New(t)
	require.Nil(err)

	require.EqualValues(http.StatusNotFound, res.StatusCode)
}

func TestMember_Get_Bad_Request_Limit(t *testing.T) {
	handler := resource.NewHandler()
	request := NewMemberRequest("",
		http.MethodGet,
		map[string]string{
			resource.CommentOrganizationName: "xendit",
		},
		map[string]string{
			"limit": "a",
		})
	res, err := api.NewHandler(handler)(GetContext(), request)

	require := require.New(t)
	require.Nil(err)

	require.EqualValues(http.StatusBadRequest, res.StatusCode)
}

func TestMember_Get_Bad_Request_Offset(t *testing.T) {
	handler := resource.NewHandler()
	request := NewMemberRequest("",
		http.MethodGet,
		map[string]string{
			resource.CommentOrganizationName: "xendit",
		},
		map[string]string{
			"offset": "a",
		})
	res, err := api.NewHandler(handler)(GetContext(), request)

	require := require.New(t)
	require.Nil(err)

	require.EqualValues(http.StatusBadRequest, res.StatusCode)
}

func TestMember_Get(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.github.com/orgs/xendit/members",
		httpmock.NewStringResponder(200, string(ReadFile("../testdata/member.json"))))

	handler := resource.NewHandler()
	request := NewMemberRequest("",
		http.MethodGet,
		map[string]string{
			resource.CommentOrganizationName: "xendit",
		},
		map[string]string{})
	res, err := api.NewHandler(handler)(GetContext(), request)

	require := require.New(t)
	require.Nil(err)

	require.EqualValues(http.StatusOK, res.StatusCode)
}
