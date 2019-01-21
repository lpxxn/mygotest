package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

const (
	sqsUrl = ""
)

var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

var svc = sqs.New(sess)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer
	fmt.Println("----begin hello------")
	out, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(fmt.Sprintf("current time %d", time.Now().Unix())),
		QueueUrl:     aws.String(sqsUrl),
	})
	outReq, err2 := svc.GetQueueAttributes(&sqs.GetQueueAttributesInput{
		AttributeNames: []*string{aws.String(sqs.QueueAttributeNameApproximateNumberOfMessages)},
		QueueUrl:       aws.String(sqsUrl),
	})
	var body []byte
	if err != nil && err2 != nil {
		body, _ = json.Marshal(map[string]interface{}{
			"message": "send sqs msg error" + err.Error(),
		})
	} else {
		//nums, _ := strconv.ParseInt(*outReq.Attributes[sqs.QueueAttributeNameApproximateNumberOfMessages], 10, 64)
		body, _ = json.Marshal(map[string]interface{}{
			"message": "send sqs executed successfully! msgId: " +
				*out.MessageId + " count of number msg: " + *outReq.Attributes[sqs.QueueAttributeNameApproximateNumberOfMessages],
		})
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	fmt.Println("-------end hello---------")
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
