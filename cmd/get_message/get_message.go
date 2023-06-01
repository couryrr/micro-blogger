package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, requestProxy events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	messageId := requestProxy.PathParameters["message_id"]
	message := fmt.Sprintf("Message %s has been requested!", messageId)
	return events.APIGatewayProxyResponse{
		Body:       message,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
