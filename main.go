package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/frozentech/api"
	"github.com/frozentech/comment/database"
	"github.com/frozentech/comment/resource"
	"github.com/frozentech/logs"
)

// Init ...
func Init() {
	var err error

	database.Connection, err = database.Connect()
	if err != nil {
		panic(err)
	}

}

func main() {
	lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, err error) {
		logs.Story = logs.New()
		resp, err = api.NewHandler(resource.NewHandler())(ctx, req)
		return
	})

	defer database.Connection.Close(context.Background())
}
