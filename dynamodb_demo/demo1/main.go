package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mygotest/dynamodb_demo/dynamodb_utils"
)
var DB *dynamodb.DynamoDB
func main() {
	svc, dbs := dynamodb_utils.StartSession()
	defer dbs.Stop()
	if err := dynamodb_utils.CreateTable(svc); err != nil {
		panic(err)
	}

	// describe table to test connection
	p := &dynamodb.DescribeTableInput{
		TableName: aws.String(dynamodb_utils.TestTable1Name), // Required
	}
	_, err := svc.DescribeTable(p)
	if err != nil {
		panic(err)
	}
	DB = svc
	p1 := &dynamodb_utils.Table1DataInfo{Key: "1", Sky1: "k1", Sky2: "k2", Sky3: "k3", Name: "li"}
	_, err = dynamodb_utils.Put(DB, p1)
	if err != nil {
		panic(err)
	}

	p2 := &dynamodb_utils.Table1DataInfo{Key: "2", Sky1: "2k1", Sky2: "2k2", Sky3: "2k3", Name: "peng"}
	_, err = dynamodb_utils.Put(DB, p2)
	if err != nil {
		panic(err)
	}
	dis := make([]dynamodb_utils.Table1DataInfo, 0)
	err = dynamodb_utils.QueryBySkey(DB, "2k1", "2k2", &dis)
	if err != nil {
		panic(err)
	}

}

