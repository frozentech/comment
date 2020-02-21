package resource_test

import (
	"net/http"
	"testing"

	"github.com/frozentech/api"
	"github.com/frozentech/comment/resource"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestComment_Delete_Access_Denied(t *testing.T) {
	handler := resource.NewHandler()
	request := NewCommentRequest("", http.MethodDelete, map[string]string{resource.CommentOrganizationName: ""}, map[string]string{})
	res, err := api.NewHandler(handler)(GetContext(), request)

	require := require.New(t)
	require.Nil(err)

	require.EqualValues(http.StatusForbidden, res.StatusCode)
}

func TestComment_Delete_Bad_Request_Organization_Error(t *testing.T) {
	handler := resource.NewHandler()
	request := NewCommentRequest("", http.MethodDelete, map[string]string{resource.CommentOrganizationName: ""}, map[string]string{})
	request.Headers = map[string]string{
		"Content-Type":             "application/json",
		resource.GithubAccessToken: "xxxxxxx",
	}
	res, err := api.NewHandler(handler)(GetContext(), request)

	require := require.New(t)
	require.Nil(err)

	require.EqualValues(http.StatusNotFound, res.StatusCode)
}

func TestComment_Delete_Bad_Request_Organization_Empty(t *testing.T) {
	handler := resource.NewHandler()
	request := NewCommentRequest("", http.MethodDelete, map[string]string{resource.CommentOrganizationName: "xendit"}, map[string]string{})
	request.Headers = map[string]string{
		"Content-Type":             "application/json",
		resource.GithubAccessToken: "xxxxxxx",
	}
	res, err := api.NewHandler(handler)(GetContext(), request)

	require := require.New(t)
	require.Nil(err)

	require.EqualValues(http.StatusNotFound, res.StatusCode)
}

func TestComment_Delete(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.github.com/orgs/xendit",
		httpmock.NewStringResponder(200, string(ReadFile("../testdata/organization.json"))))

	httpmock.RegisterResponder("GET", "https://api.github.com/user",
		httpmock.NewStringResponder(400, string(ReadFile("../testdata/access_denied.json"))))

	handler := resource.NewHandler()
	request := NewCommentRequest(``, http.MethodDelete, map[string]string{resource.CommentOrganizationName: "xendit"}, map[string]string{})
	request.Headers = map[string]string{
		"Content-Type":             "application/json",
		resource.GithubAccessToken: "xxxxxxx",
	}
	res, err := api.NewHandler(handler)(GetContext(), request)

	require := require.New(t)
	require.Nil(err)

	require.EqualValues(http.StatusNoContent, res.StatusCode)
}
