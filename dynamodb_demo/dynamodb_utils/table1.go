package dynamodb_utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type testTabel1 struct {
	Dao *dynamodb.DynamoDB
	TableName string
}


var TestTable1Dao *testTabel1

func setTestTable1Dao(db *dynamodb.DynamoDB, tableName string) {
	TestTable1Dao = &testTabel1{
		Dao:db,
		TableName:tableName,
	}
}

const (
	TestTable1Name = "testTable1Name"
	Table1KVPrimaryKey = "key"
	Table1KVSecondaryKey1 = "skey1"
	Table1KVSecondaryKey2 = "skey2"
	Table1KVSecondaryKey3 = "skey3"
)

type Table1DataInfo struct {
	Key string `json:"key"`
}

var Test1TableDescription = dynamodb.TableDescription{
	AttributeDefinitions: []*dynamodb.AttributeDefinition {
		{
			AttributeName: aws.String(Table1KVPrimaryKey),
			AttributeType: aws.String("S"),
		},
		{
			AttributeName: aws.String(Table1KVPrimaryKey),
			AttributeType: aws.String("S"),
		},
	},
	GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndexDescription{
		{
			IndexName: aws.String(Table1KVSecondaryKey1),
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String(Table1KVSecondaryKey1),
					KeyType:       aws.String("HASH"),
				},
			},
			Projection: &dynamodb.Projection{
				ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
			},
		},
		{
			IndexName: aws.String(Table1KVSecondaryKey2),
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String(Table1KVSecondaryKey2),
					KeyType:       aws.String("HASH"),
				},
			},
			Projection: &dynamodb.Projection{
				ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
			},
		},
		{
			IndexName: aws.String(Table1KVSecondaryKey3),
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String(Table1KVSecondaryKey3),
					KeyType:       aws.String("HASH"),
				},
			},
			Projection: &dynamodb.Projection{
				ProjectionType: aws.String(dynamodb.ProjectionTypeAll),

			},
		},
	},
	KeySchema: []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String(Table1KVPrimaryKey),
			KeyType:       aws.String("HASH"),
		},
	},
}
