package database_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/frozentech/comment/database"
	"github.com/frozentech/comment/model"
	"github.com/stretchr/testify/assert"
)

func TestDatabase_Connect(t *testing.T) {
	ctx := context.Background()
	assert := assert.New(t)
	server := os.Getenv("DATABASE_URL")
	defer func() {
		os.Setenv("DATABASE_URL", server)
	}()

	os.Setenv("DATABASE_URL", "sss")
	_, err := database.Connect()
	assert.Error(err)

	// THROW AWAY SERVER
	os.Setenv("DATABASE_URL", "postgres://nbubghfs:Xr71_4P1hz8sF0Ia_l9yY8GI07VuIQiJ@arjuna.db.elephantsql.com:5432/nbubghfs")
	conn, err := database.Connect()
	assert.Nil(err)

	conn.Exec(ctx, database.SQLDropTableComment)

	_, err = conn.Exec(ctx, database.SQLCreateTableComment)
	assert.Nil(err)

	database.Connection = conn

	for i := 0; i < 10; i++ {
		comment := model.NewComment()
		comment.Member = model.Member{
			ID:      int64(1),
			Name:    "Name",
			Profile: fmt.Sprintf("Profile %d", i),
		}
		comment.Organization = "xendit"
		comment.Message = fmt.Sprintf("message %d", i)
		comment.Prepare()

		err = database.PostComment(ctx, *comment)
		assert.Nil(err)
	}

	comments, err := database.GetComment(ctx, "xendit")
	assert.Nil(err)
	assert.True(len(comments) > 0)

	err = database.DeleteComment(ctx, "xendit")
	assert.Nil(err)

	_, err = database.GetComment(ctx, "xendit")
	assert.Error(err)
}
