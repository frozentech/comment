package model_test

import (
	"testing"

	"github.com/frozentech/comment/model"
	"github.com/stretchr/testify/assert"
)

func TestModel_GenerateUUID_V4(t *testing.T) {
	id := model.GenerateUUID()
	assert := assert.New(t)
	assert.NotEmpty(id)
}

func TestModel_GenerateUUID_V3(t *testing.T) {
	id := model.GenerateUUID("id")
	assert := assert.New(t)
	assert.NotEmpty(id)
}
