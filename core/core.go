package core

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const (
	region    = "us-west-2"
	tableName = "heavybee-table"
)

type Item struct {
	FromName   string `json:"fromName" dynamodbav:"fromName"`
	FromEmail  string `json:"fromEmail" dynamodbav:"fromEmail"`
	FromDomain string `json:"fromDomain" dynamodbav:"fromDomain"`
}

func RunBulkInsert(jsonFile string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	file, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	var items []Item
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&items)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for _, item := range items {
		av, err := attributevalue.MarshalMap(item)
		if err != nil {
			fmt.Println("Error marshalling item:", err)
			return
		}
		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(tableName),
		}
		_, err = svc.PutItem(context.TODO(), input)
		if err != nil {
			fmt.Println("Error putting item:", err)
			return
		}
	}
	fmt.Printf("Inserted %d items successfully\n", len(items))
}

func QueryItems() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := svc.Scan(context.TODO(), input)
	if err != nil {
		fmt.Println("Error scanning table:", err)
		return
	}

	var items []Item
	err = attributevalue.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		fmt.Println("Error unmarshalling items:", err)
		return
	}

	fmt.Println("Queried items:")
	for _, item := range items {
		fmt.Printf("FromName: %s, FromEmail: %s, FromDomain: %s\n", item.FromName, item.FromEmail, item.FromDomain)
	}
}
