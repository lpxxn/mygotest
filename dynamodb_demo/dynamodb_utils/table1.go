package dynamodb_utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/pkg/errors"
)

type testTabel1 struct {
	Dao       *dynamodb.DynamoDB
	TableName string
}

var TestTable1Dao *testTabel1

func setTestTable1Dao(db *dynamodb.DynamoDB, tableName string) {
	TestTable1Dao = &testTabel1{
		Dao:       db,
		TableName: tableName,
	}
}

const (
	TestTable1Name        = "testTable1Name"
	Table1KVPrimaryKey    = "key"
	Table1KvPrimaryRange  = "type"
	Table1KVSecondaryKey1 = "skey1"
	Table1KVSecondaryKey2 = "skey2"
	Table1KVSecondaryKey3 = "skey3"
)

type Table1DataInfo struct {
	Key  string `json:"key" dynamodbav:"key"`
	Type string `json:"type" dynamodbav:"type"`
	Sky1 string `json:"sky1" dynamodbav:"skey1"`
	Sky2 string `json:"sky2" dynamodbav:"skey2"`
	Sky3 string `json:"sky3" dynamodbav:"skey3"`
	Name string `json:"name" dynamodbav:"name"`
}

var Test1TableDescription = dynamodb.TableDescription{
	TableName: aws.String(TestTable1Name),
	AttributeDefinitions: []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String(Table1KVPrimaryKey),
			AttributeType: aws.String("S"),
		},
		{
			AttributeName: aws.String(Table1KvPrimaryRange),
			AttributeType: aws.String("S"),
		},
		{
			AttributeName: aws.String(Table1KVSecondaryKey1),
			AttributeType: aws.String("S"),
		},
		{
			AttributeName: aws.String(Table1KVSecondaryKey2),
			AttributeType: aws.String("S"),
		},
		{
			AttributeName: aws.String(Table1KVSecondaryKey3),
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
		{
			AttributeName: aws.String(Table1KvPrimaryRange),
			KeyType:       aws.String("RANGE"),
		},
	},
}

func RetrieveTable1GlobalSecondaryIndexDescription(idxName string) *dynamodb.GlobalSecondaryIndexDescription {
	for i := range Test1TableDescription.GlobalSecondaryIndexes {
		g2idx := Test1TableDescription.GlobalSecondaryIndexes[i]
		if aws.StringValue(g2idx.KeySchema[0].AttributeName) == idxName {
			return g2idx
		}
	}
	return nil
}

func RetrieveGlobalSecondaryIndexDefinition(idxName string) *dynamodb.AttributeDefinition {
	for i := range Test1TableDescription.AttributeDefinitions {
		def := Test1TableDescription.AttributeDefinitions[i]
		if aws.StringValue(def.AttributeName) == idxName {
			return def
		}
	}
	return Test1TableDescription.AttributeDefinitions[1]
}

func MakeAttributeValue(definition *dynamodb.AttributeDefinition, keyvalue string) *dynamodb.AttributeValue {
	switch aws.StringValue(definition.AttributeType) {
	case dynamodb.ScalarAttributeTypeN:
		return &dynamodb.AttributeValue{
			N: aws.String(keyvalue),
		}
	case dynamodb.ScalarAttributeTypeB:
		return &dynamodb.AttributeValue{
			B: []byte(keyvalue),
		}
	}
	return &dynamodb.AttributeValue{
		S: aws.String(keyvalue),
	}
}

var ErrNotFound = errors.New("record not found")

const (
	expressionAttributeNameKey1  = "#key1"
	expressionAttributeValueKey1 = ":key1"

	expressionAttributeNameKey2  = "#key2"
	expressionAttributeValueKey2 = ":key2"
)

func Put(db *dynamodb.DynamoDB, o *Table1DataInfo) (*dynamodb.PutItemOutput, error) {
	m, err := dynamodbattribute.MarshalMap(o)
	if err != nil {
		return nil, errors.WithMessage(err, "PlanetKVDao: dynamodbattribute.MarshalMap")
	}
	in := &dynamodb.PutItemInput{
		Item:      m,
		TableName: aws.String(TestTable1Name),
	}
	return db.PutItem(in)
}

func QueryBySkey(db *dynamodb.DynamoDB, idx1Val string, idx2Val string, dist interface{}) error {
	idx1Name := Table1KVSecondaryKey1
	g2Key1 := RetrieveTable1GlobalSecondaryIndexDescription(idx1Name)
	idx2Name := Table1KVSecondaryKey2
	g2Key2 := RetrieveTable1GlobalSecondaryIndexDescription(idx2Name)

	/// query 只能查询一个指定的分区键
	qInput := &dynamodb.QueryInput{
		IndexName: aws.String(idx1Name),
		//KeyConditionExpression: aws.String(fmt.Sprintf("%s = %s and %s = %s", expressionAttributeNameKey1, expressionAttributeValueKey1, expressionAttributeNameKey2, expressionAttributeValueKey2)),
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = %s", expressionAttributeNameKey1, expressionAttributeValueKey1)),
		FilterExpression:       aws.String(fmt.Sprintf("%s = %s", expressionAttributeNameKey2, expressionAttributeValueKey2)),
		ExpressionAttributeNames: map[string]*string{
			expressionAttributeNameKey1: g2Key1.KeySchema[0].AttributeName,
			expressionAttributeNameKey2: g2Key2.KeySchema[0].AttributeName,
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			expressionAttributeValueKey1: MakeAttributeValue(RetrieveGlobalSecondaryIndexDefinition(idx1Name), idx1Val),
			expressionAttributeValueKey2: MakeAttributeValue(RetrieveGlobalSecondaryIndexDefinition(idx2Name), idx2Val),
		},
		TableName: aws.String(TestTable1Name),
	}
	result, err := db.Query(qInput)
	_ = result
	//
	sInput := &dynamodb.ScanInput{
		FilterExpression: aws.String(fmt.Sprintf("%s = %s and %s = %s", expressionAttributeNameKey1, expressionAttributeValueKey1, expressionAttributeNameKey2, expressionAttributeValueKey2)),
		ExpressionAttributeNames: map[string]*string{
			expressionAttributeNameKey1: g2Key1.KeySchema[0].AttributeName,
			expressionAttributeNameKey2: g2Key2.KeySchema[0].AttributeName,
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			expressionAttributeValueKey1: MakeAttributeValue(RetrieveGlobalSecondaryIndexDefinition(idx1Name), idx1Val),
			expressionAttributeValueKey2: MakeAttributeValue(RetrieveGlobalSecondaryIndexDefinition(idx2Name), idx2Val),
		},
		TableName: aws.String(TestTable1Name),
	}
	sResult, err := db.Scan(sInput)
	if err != nil {
		return err
	}
	return dynamodbattribute.UnmarshalListOfMaps(sResult.Items, dist)
}

func primaryPartitionKeyDefinition() *dynamodb.AttributeDefinition {
	pkname := aws.String(PrimaryPartitionKeyName())
	for i := range Test1TableDescription.AttributeDefinitions {
		def := Test1TableDescription.AttributeDefinitions[i]
		if aws.String(*def.AttributeName) == pkname {
			return def
		}
	}
	return Test1TableDescription.AttributeDefinitions[0]
}

func makePrimaryPartitionKeyAttributeValue(keyvalue string) *dynamodb.AttributeValue {
	return MakeAttributeValue(primaryPartitionKeyDefinition(), keyvalue)
}

func PrimaryPartitionKeyName() string {
	return *Test1TableDescription.KeySchema[0].AttributeName
}

// GetOption options for Get function
type GetOption func(input *dynamodb.GetItemInput)

func GetItemByKey(db *dynamodb.DynamoDB, keyValue string, dist interface{}, opts ...GetOption) error {
	gInput := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			PrimaryPartitionKeyName(): makePrimaryPartitionKeyAttributeValue(keyValue),
		},
		TableName: Test1TableDescription.TableName,
	}
	for _, opt := range opts {
		opt(gInput)
	}
	out, err := db.GetItem(gInput)
	if err != nil {
		return err
	}
	if len(out.Item) == 0 {
		return ErrNotFound
	}

	return dynamodbattribute.UnmarshalMap(out.Item, dist)
}
