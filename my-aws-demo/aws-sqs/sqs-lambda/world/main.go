package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	if len(sqsEvent.Records) == 0 {
		return errors.New("No SQS message passed to function")
	}
	fmt.Println("-------lipeng---------")
	for _, msg := range sqsEvent.Records {
		fmt.Printf("Got SQS message %q with body %q\n", msg.MessageId, msg.Body)
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
