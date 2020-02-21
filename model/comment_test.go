package model_test

import (
	"testing"

	"github.com/frozentech/comment/model"
	"github.com/stretchr/testify/assert"
)

func TestComment_NewComment(t *testing.T) {
	comment := model.NewComment()

	assert := assert.New(t)
	assert.IsType(&model.Comment{}, comment)
}

func TestComment_Prepare(t *testing.T) {
	comment := model.NewComment()
	comment.Member = model.Member{
		ID:      int64(1),
		Name:    "Name",
		Profile: "Profile",
	}
	comment.Prepare()

	assert := assert.New(t)
	assert.Equal(comment.MemberAsString, `{"id":1,"login":"Name","avatar_url":"Profile"}`)
}

func TestComment_Format(t *testing.T) {
	comment := model.NewComment()
	comment.MemberAsString = `{"id":"1","login":"Name","avatar_url":"Profile"}`
	comment.Format()

	assert := assert.New(t)
	assert.Equal(comment.Member.ID, comment.Member.ID)
	assert.Equal(comment.Member.Name, comment.Member.Name)
	assert.Equal(comment.Member.Profile, comment.Member.Profile)
}
