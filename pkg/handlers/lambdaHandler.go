package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/satvirgrewal/lambda-go/pkg/user"
)

type Handler interface {
	GetUser(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	CreateUser(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

type LambdaHandler struct {
}

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func NewLambdaHandler() Handler {
	return &LambdaHandler{}
}

func (handler *LambdaHandler) GetUser(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		// Get single user
		result, err := user.FetchUser(email)
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return apiResponse(http.StatusOK, result)
	}
	return apiResponse(http.StatusBadRequest, ErrorBody{
		aws.String("email not provided"),
	})
}

func (handler *LambdaHandler) CreateUser(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	result, err := user.CreateUser(req)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusCreated, result)
}

func apiResponse(status int, body interface{}) (events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return resp, nil
}
