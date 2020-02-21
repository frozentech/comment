package resource_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/frozentech/comment/database"
	"github.com/frozentech/comment/model"
	"github.com/frozentech/logs"
)

func CleanUpTable() {
	ctx := context.Background()
	database.Connection.Exec(ctx, database.SQLDropTableComment)
	database.Connection.Exec(ctx, database.SQLCreateTableComment)
}

func FillUpTable() {
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

		database.PostComment(context.Background(), *comment)
	}
}

func InitDatabase() {

	CleanUpTable()
}

func Init() {
	logs.Story = logs.New()

	server := os.Getenv("DATABASE_URL")
	defer func() {
		os.Setenv("DATABASE_URL", server)
	}()

	os.Setenv("DATABASE_URL", "postgres://nbubghfs:Xr71_4P1hz8sF0Ia_l9yY8GI07VuIQiJ@arjuna.db.elephantsql.com:5432/nbubghfs")
	conn, _ := database.Connect()
	database.Connection = conn

	InitDatabase()
}

func TestMain(m *testing.M) {
	Init()
	os.Exit(m.Run())
}

func GetContext() context.Context {
	return context.WithValue(context.Background(), "x-amzn-trace-id", "Root=fakeid; Parent=reqid; Sampled=1")
}

// ReadFile ...
func ReadFile(path string) []byte {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return body
}
