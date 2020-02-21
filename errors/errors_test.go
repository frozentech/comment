package errors_test

import (
	"testing"

	"github.com/frozentech/comment/errors"
	"github.com/stretchr/testify/assert"
)

func Test_ListErrors(t *testing.T) {
	list := errors.ListErrors()
	assert := assert.New(t)
	assert.NotEmpty(list)
}

func Test_New(t *testing.T) {
	list := errors.New(errors.ErrorResourceNotFound, "body")
	assert := assert.New(t)
	assert.NotEmpty(list)

	list = errors.New(errors.ErrorResourceNotFound, "")
	assert.NotEmpty(list)

	list = errors.New("", "")
	assert.NotEmpty(list)
}

func Test_NewMessage(t *testing.T) {
	assert := assert.New(t)

	list := errors.NewMessage(errors.ErrorResourceNotFound, "body")
	assert.NotEmpty(list)

	list = errors.NewMessage(errors.ErrorInvalidRequestBody, "body")
	assert.NotEmpty(list)

	list = errors.NewMessage(errors.ErrorInvalidRangeValue, "body")
	assert.NotEmpty(list)

	list = errors.NewMessage(errors.ErrorEmptyRequestBody, "body")
	assert.NotEmpty(list)

	list = errors.NewMessage(errors.ErrorAccessDenied, "body")
	assert.NotEmpty(list)

	list = errors.NewMessage("", "")
	assert.NotEmpty(list)

}
