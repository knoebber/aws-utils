package lambda

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Start calls github.com/aws/aws-lambda-go/lambda.Start
func Start(handler interface{}) {
	lambda.Start(handler)
}

// Start calls github.com/aws/aws-lambda-go/lambda.StartHandler
func StartHandler(handler lambda.Handler) {
	lambda.StartHandler(handler)
}

// SetResponseMessage sets an APIGatewayProxyResponse body like {"message": "..."}.
func SetResponseMessage(response *events.APIGatewayProxyResponse, message string) {
	SetResponseBody(response, map[string]string{"message": message})
}

// SetResponseBody marshals body into an APIGatewayProxyResponse body.
func SetResponseBody(response *events.APIGatewayProxyResponse, responseBody interface{}) {
	bytes, err := json.Marshal(responseBody)
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		return
	}

	response.Body = string(bytes)
	response.StatusCode = http.StatusOK
}

// Env returns the value for an enviroment variable or an error when it doesn't exist.
func Env(key string) (value string, err error) {
	if value = os.Getenv(key); value == "" {
		err = fmt.Errorf("$%s is not set or empty", key)
	}
	return
}
