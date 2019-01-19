package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"
	"os/signal"
	"time"
)

type Sqs struct {
	Sqs *sqs.SQS
	Url string
}

func (s *Sqs) SendMsg(body string) (string, error) {
	out, err := s.Sqs.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(body),
		QueueUrl:    aws.String(s.Url),
	})
	if err != nil {
		return "", err
	}
	return *out.MessageId, nil
}

func (s *Sqs) ReceiveMsg(max int64) ([]*sqs.Message, error) {
	out, err := s.Sqs.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:              aws.String(s.Url),
		MaxNumberOfMessages:   aws.Int64(max),
		MessageAttributeNames: []*string{aws.String("All")},
	})
	return out.Messages, err
}

func (s *Sqs) DelMsg(receipt *string) error {
	_, err := s.Sqs.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(s.Url),
		ReceiptHandle: receipt,
	})
	return err
}

const (
	endPoint = ""
)

func main() {
	newSession, _ := session.NewSession(&aws.Config{
		Endpoint: aws.String(endPoint),
		Region:   aws.String("cn-northwest-1"),
	})
	svc := sqs.New(newSession)
	sqs := new(Sqs)
	sqs.Sqs = svc
	sqs.Url = endPoint

	consumerCh := time.Tick(time.Second * 5)
	go func() {
		for {
			<-consumerCh
			outMsg, err := sqs.ReceiveMsg(10)
			if err != nil {
				fmt.Println("have error :", err)
			} else {
				fmt.Println("len of Message: ", len(outMsg))
				fmt.Printf("out Message :%#v \n --------\n", outMsg)
				for _, v := range outMsg {
					err := sqs.DelMsg(v.ReceiptHandle)
					fmt.Println("delete err :", err)
				}
			}
		}
	}()

	productorCh := time.Tick(time.Second * 3)
	go func() {
		for {
			pt := <-productorCh
			t := pt.String()
			outID, err := sqs.SendMsg("Hello world t: " + t)
			if err != nil {
				fmt.Println("have error :", err)
			} else {
				fmt.Println("outID: ", outID)
			}
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	fmt.Println("stop server....")
}
