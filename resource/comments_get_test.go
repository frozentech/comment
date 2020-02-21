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

func NewTestComment() resource.Comment {
	return resource.NewComment()
}

func NewCommentRequest(payload string, method string, path map[string]string, query map[string]string) events.APIGatewayProxyRequest {
	comment := NewTestComment()

	return events.APIGatewayProxyRequest{
		Resource: comment.URI,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		PathParameters:        path,
		QueryStringParameters: query,
		HTTPMethod:            method,
		Body:                  payload,
	}
}

func TestComment_New(t *testing.T) {
	a := NewTestComment()
	assert := assert.New(t)
	assert.IsType(resource.Comment{}, a)
	assert.Equal("orgs/{organization}/comments", a.URI)
}

func TestComment_Methods(t *testing.T) {
	a := NewTestComment()
	assert := assert.New(t)
	assert.Equal(a.AllowedMethods(), []string{http.MethodGet, http.MethodPost, http.MethodDelete})
}

func TestComment_Get_Not_Found(t *testing.T) {
	handler := resource.NewHandler()
	request := NewCommentRequest("", http.MethodGet, map[string]string{resource.CommentOrganizationName: ""}, map[string]string{})
	res, err := api.NewHandler(handler)(GetContext(), request)

	require := require.New(t)
	require.Nil(err)

	require.EqualValues(http.StatusNotFound, res.StatusCode)
}

func TestComment_Get_Bad_Request_Organization_Error(t *testing.T) {
	handler := resource.NewHandler()
	request := NewCommentRequest("", http.MethodGet, map[string]string{resource.CommentOrganizationName: "xendit"}, map[string]string{})
	res, err := api.NewHandler(handler)(GetContext(), request)

	require := require.New(t)
	require.Nil(err)

	require.EqualValues(http.StatusNotFound, res.StatusCode)
}

func TestComment_Get_Bad_Request_Comments_Error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.github.com/orgs/xendit",
		httpmock.NewStringResponder(200, string(ReadFile("../testdata/organization.json"))))

	handler := resource.NewHandler()
	request := NewCommentRequest("", http.MethodGet, map[string]string{resource.CommentOrganizationName: "xendit"}, map[string]string{})
	res, err := api.NewHandler(handler)(GetContext(), request)

	require := require.New(t)
	require.Nil(err)

	require.EqualValues(http.StatusNotFound, res.StatusCode)
}

func TestComment_Get_Bad_Request_Comments_Found(t *testing.T) {
	CleanUpTable()
	FillUpTable()

	httpmock.Activate()
	defer func() {
		httpmock.DeactivateAndReset()
	}()

	httpmock.RegisterResponder("GET", "https://api.github.com/orgs/xendit",
		httpmock.NewStringResponder(200, string(ReadFile("../testdata/organization.json"))))

	handler := resource.NewHandler()
	request := NewCommentRequest("", http.MethodGet, map[string]string{resource.CommentOrganizationName: "xendit"}, map[string]string{})
	res, err := api.NewHandler(handler)(GetContext(), request)

	require := require.New(t)
	require.Nil(err)

	require.EqualValues(http.StatusOK, res.StatusCode)

}
