package dynamodb_utils

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func StartSession() (*dynamodb.DynamoDB, *server.DynamoDBServer) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	accessSecret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	s := server.Start(accessKey, accessSecret, "48080")
	theCredentials := credentials.NewStaticCredentials(accessKey, accessSecret, "")

	sess := session.Must(session.NewSession())
	awsc := &aws.Config{
		Credentials: theCredentials,
		Region:      aws.String("cn-north-1"),
	}
	awsc.Endpoint = aws.String("http://127.0.0.1:48080")
	return dynamodb.New(sess, awsc), s
}
