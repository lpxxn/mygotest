package dynamodb_utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go.planetmeican.com/planet/dynamodb-local-server"
)

func StartSession() (*dynamodb.DynamoDB, *server.DynamoDBServer) {
	s :=server.Start("AKIAPA7LKDMYJU2P2OVA", "BRn5WzmIShvCCvRXchHVDTacX8ZS9FO/d3uleo2C", "48080")
	theCredentials := credentials.NewStaticCredentials("AKIAPA7LKDMYJU2P2OVA", "BRn5WzmIShvCCvRXchHVDTacX8ZS9FO/d3uleo2C", "")

	sess := session.Must(session.NewSession())
	awsc := &aws.Config{
		Credentials: theCredentials,
		Region:      aws.String("cn-north-1"),
	}
	awsc.Endpoint = aws.String("http://127.0.0.1:48080")
	return dynamodb.New(sess, awsc), s
}