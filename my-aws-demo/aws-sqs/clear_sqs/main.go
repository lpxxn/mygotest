package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	endPoint string
	region   string
)

type Sqs struct {
	Sqs *sqs.SQS
	Url string
}

func (s *Sqs) DelAllMsg() error {
	_, err := s.Sqs.PurgeQueue(&sqs.PurgeQueueInput{
		QueueUrl: aws.String(s.Url),
	})
	return err
}

func init() {
	flag.StringVar(&endPoint, "endPoint", "", "endPoint")
	flag.StringVar(&region, "region", "cn-northwest-1", "region")
	flag.Parse()
}

func main() {
	if endPoint == "" {
		fmt.Println("no endpoint")
		return
	}
	newSession, _ := session.NewSession(&aws.Config{
		// Endpoint: aws.String(endPoint),
		Region: aws.String(region),
	})
	svc := sqs.New(newSession)
	sqs := new(Sqs)
	v, err := newSession.Config.Credentials.Get()
	fmt.Println(err, v.AccessKeyID, " ", v.SecretAccessKey)
	sqs.Sqs = svc
	sqs.Url = endPoint

	if err := sqs.DelAllMsg(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("clear successfully")
	}
}
