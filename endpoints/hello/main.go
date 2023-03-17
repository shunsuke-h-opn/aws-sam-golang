package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"os"
)

var initialized = false
var ginLambda *ginadapter.GinLambda

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v, %v", os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_USER")),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
