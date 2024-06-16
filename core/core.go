package core

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Item struct {
	FromName   string `dynamodbav:"fromName"`
	FromEmail  string `dynamodbav:"fromEmail"`
	FromDomain string `dynamodbav:"fromDomain"`
}

func Run() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	item := Item{
		FromName:   "John Doe",
		FromEmail:  "johndoe@example.com",
		FromDomain: "example.com",
	}
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		fmt.Println("Error marshalling item:", err)
		return
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("heavybee-table"),
	}
	_, err = svc.PutItem(context.TODO(), input)
	if err != nil {
		fmt.Println("Error putting item:", err)
		return
	}
	fmt.Println("Item created successfully")
}
