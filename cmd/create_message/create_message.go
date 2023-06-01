package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type CreateMessage struct {
	Id      string `json:"id" dynamodbav:"id"`
	Title   string `json:"title" dynamodbav:"title"`
	Content string `json:"content" dynamodbav:"content"`
	Author  string `json:"author" dynamodbav:"author"`
}

func HandleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-2"
		return nil
	})

	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	for _, message := range sqsEvent.Records {
		var createMessage CreateMessage

		err = json.Unmarshal([]byte(message.Body), &createMessage)

		createMessage.Id = message.MessageId

		if err != nil {
			panic(err)
		}

		item, err := attributevalue.MarshalMap(createMessage)

		fmt.Printf("marshalled struct: %+v", item)
		fmt.Println()

		if err != nil {
			panic(err)
		}

		var created *dynamodb.PutItemOutput

		created, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
			TableName: aws.String("BlogEntries"), Item: item,
		})

		if err != nil {
			panic(err)
		}

		fmt.Println("Start log")
		fmt.Println(created.Attributes)
		fmt.Println("End log")
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
