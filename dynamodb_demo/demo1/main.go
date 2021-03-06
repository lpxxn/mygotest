package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mygotest/dynamodb_demo/dynamodb_utils"

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
	p0 := &dynamodb_utils.Table1DataInfo{Key: "0", Sky1: "k1", Sky2: "k2", Sky3: "k3", Name: "li", Type: "t1"}
	_, err = dynamodb_utils.Put(DB, p0)
	if err != nil {
		panic(err)
	}

	p1 := &dynamodb_utils.Table1DataInfo{Key: "1", Sky1: "k1", Sky2: "k2", Sky3: "k4", Name: "li", Type: "t1"}
	_, err = dynamodb_utils.Put(DB, p1)
	if err != nil {
		panic(err)
	}

	p2 := &dynamodb_utils.Table1DataInfo{Key: "2", Sky1: "2k1", Sky2: "2k2", Sky3: "2k3", Name: "peng", Type: "t1"}
	_, err = dynamodb_utils.Put(DB, p2)
	if err != nil {
		panic(err)
	}
	p3 := &dynamodb_utils.Table1DataInfo{Key: "3", Sky1: "3k1", Sky2: "3k2", Sky3: "3k3", Name: "Abc", Type: "t2"}

		_, err = dynamodb_utils.Put(DB, p3)
	if err != nil {
		panic(err)
	}
	dis := make([]dynamodb_utils.Table1DataInfo, 0)
	err = dynamodb_utils.QueryBySkey(DB, "2k1", "2k2", &dis)
	if err != nil {
		panic(err)
	}

	dis2 := make([]dynamodb_utils.Table1DataInfo, 0)
	err = dynamodb_utils.QueryBySkey2(DB, "2k1", "t1", &dis2)
	if err != nil {
		fmt.Println(err)
	}

	dis3 := make([]dynamodb_utils.Table1DataInfo, 0)
	err = dynamodb_utils.QueryBySkey3(DB, "k1", "", &dis3)
	if err != nil {
		fmt.Println(err)
	}

	disItem := new(dynamodb_utils.Table1DataInfo)
	// 方法 GetItem 定义中有range 在查询的时候就要给上。secondary 有range也可可以不用传range
	f := func(input *dynamodb.GetItemInput) {
		input.Key[dynamodb_utils.Table1KvPrimaryRange] = &dynamodb.AttributeValue{
			S: aws.String("t2"),
		}
	}
	if err := dynamodb_utils.GetItemByKey(DB, "3", disItem, f); err != nil {
		fmt.Println(err)
	}
	if err := dynamodb_utils.GetItemByKey(DB, "2", disItem, f); err != nil {
		// not found
		fmt.Println(err)
	}
	dao, err := dynamodbdao.New(DB, dynamodb_utils.TestTable1Name, func(d *dynamodbdao.GenericDao) {

	})
	if err != nil {
		panic(err)
	}
	dis1Dao := make([]dynamodb_utils.Table1DataInfo, 0)
	err = dao.QueryByGlobalSecondaryIndex(dynamodb_utils.Table1KVSecondaryKey1, "2k1", &dis1Dao, func(input *dynamodb.QueryInput) {
		ev := fmt.Sprintf(":%s", dynamodb_utils.Table1KvPrimaryRange)
		eName := fmt.Sprintf("#%s", dynamodb_utils.Table1KvPrimaryRange)
		input.KeyConditionExpression = aws.String(fmt.Sprintf("%s and %s = %s", aws.StringValue(input.KeyConditionExpression), eName, ev))
		input.ExpressionAttributeNames[eName] = aws.String(dynamodb_utils.Table1KvPrimaryRange)
		input.ExpressionAttributeValues[ev] = &dynamodb.AttributeValue{S: aws.String("t1")}
	})
	if err != nil {
		panic(err)
	}
	dis1Dao = make([]dynamodb_utils.Table1DataInfo, 0)
	err = dao.QueryByGlobalSecondaryIndex(dynamodb_utils.Table1KVSecondaryKey1, "k1", &dis1Dao, func(input *dynamodb.QueryInput) {
		ev := fmt.Sprintf(":%s", dynamodb_utils.Table1KVSecondaryKey3)
		eName := fmt.Sprintf("#%s", dynamodb_utils.Table1KVSecondaryKey3)
		input.FilterExpression = aws.String(fmt.Sprintf("%s = %s", eName, ev))
		input.ExpressionAttributeNames[eName] = aws.String(dynamodb_utils.Table1KVSecondaryKey3)
		input.ExpressionAttributeValues[ev] = &dynamodb.AttributeValue{S: aws.String("k4")}
	})
	if err != nil {
		panic(err)
	}
}

