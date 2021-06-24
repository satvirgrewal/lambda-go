package main

import (
	"context"

	"github.com/aquasecurity/lmdrouter"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	handlers "github.com/satvirgrewal/lambda-go/pkg/handlers"
)

var router *lmdrouter.Router

const tableName = "LambdaInGoUser"

func main() {
	lambda.Start(router.Handler)
}

var handler = handlers.NewLambdaHandler()

func init() {
	router = lmdrouter.NewRouter("/users")
	router.Route("GET", "", GetUser)
	router.Route("POST", "", CreateUser)
}

func GetUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return handler.GetUser(ctx, request)
}

// Get an existing cart by id
func CreateUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return handler.CreateUser(ctx, request)
}
