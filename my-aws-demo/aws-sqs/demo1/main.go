package main

import (
	"crypto/rand"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"io"
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

func (s *Sqs) SendBatchMsg() error {
	uuid, _ := uuid()
	msg := time.Now().String()
	out, err := s.Sqs.SendMessageBatch(&sqs.SendMessageBatchInput{
		Entries: []*sqs.SendMessageBatchRequestEntry{
			&sqs.SendMessageBatchRequestEntry{
				Id: aws.String(uuid),
				MessageAttributes: map[string]*sqs.MessageAttributeValue{
					"userData": {
						DataType:    aws.String("String"),
						StringValue: aws.String(msg),
					},
				},
				MessageBody: aws.String(msg),
			},
		},
		QueueUrl: aws.String(s.Url),
	})
	if err != nil {
		return err
	}
	_ = out.String
	return err
}

func uuid() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
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
func (s *Sqs) DelAllMsg() error {
	_, err := s.Sqs.PurgeQueue(&sqs.PurgeQueueInput{
		QueueUrl: aws.String(s.Url),
	})
	return err
}

const (
	endPoint = ""
)

func main() {
	newSession, _ := session.NewSession(&aws.Config{
		// Endpoint: aws.String(endPoint),
		Region: aws.String("cn-northwest-1"),
	})
	svc := sqs.New(newSession)
	sqs := new(Sqs)
	sqs.Sqs = svc
	sqs.Url = endPoint

	if err := sqs.DelAllMsg(); err != nil {
		fmt.Println(err)
		return
	}

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
			return
			t := pt.String()
			outID, err := sqs.SendMsg("Hello world t: " + t)

			if err != nil {
				fmt.Println("have error :", err)
			} else {
				fmt.Println("outID: ", outID)
			}
			if err := sqs.SendBatchMsg(); err != nil {
				fmt.Println("SendBatchMsg have error :", err)
			}
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	fmt.Println("stop server....")
}
