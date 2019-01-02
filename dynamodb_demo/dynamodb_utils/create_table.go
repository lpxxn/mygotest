package dynamodb_utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateTable(db *dynamodb.DynamoDB) error {
	g2idx := make([]*dynamodb.GlobalSecondaryIndex, len(Test1TableDescription.GlobalSecondaryIndexes))
	for i := range Test1TableDescription.GlobalSecondaryIndexes {
		g2idx[i] = &dynamodb.GlobalSecondaryIndex{
			IndexName: Test1TableDescription.GlobalSecondaryIndexes[i].IndexName,
			KeySchema: Test1TableDescription.GlobalSecondaryIndexes[i].KeySchema,
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(10),
				WriteCapacityUnits: aws.Int64(10),
			},
			Projection: Test1TableDescription.GlobalSecondaryIndexes[i].Projection,
		}
	}
	in := &dynamodb.CreateTableInput{
		AttributeDefinitions: Test1TableDescription.AttributeDefinitions,
		KeySchema:            Test1TableDescription.KeySchema,
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		GlobalSecondaryIndexes: g2idx,
		TableName:              aws.String(TestTable1Name),
	}
	if _, err := db.CreateTable(in); err != nil {
		return err
	}
	return nil
}