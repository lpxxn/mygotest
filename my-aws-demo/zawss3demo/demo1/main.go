package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"fmt"
	"strings"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)


const (
	regin = "ap-northeast-1"
	//regin = "us-east-2"
	key_id = "AKIAIP65TRLVUOMNKALA"
	secret_key = "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"
)
var bucket = "lpxxntestbucket3"
var keyfile = "awstestkey"

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(regin),
		Credentials: credentials.NewStaticCredentials(key_id, secret_key, ""),

	})
	if err != nil {
		panic(err)
	}

	svc := s3.New(sess)

	ov, err := svc.HeadBucket(&s3.HeadBucketInput{
		Bucket:aws.String(bucket),
	})
	fmt.Println(ov.String(), " err")

	result, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		panic("Failed to list buckets")

	}

	names := map[string]struct{}{}
	for _, bucket_item := range result.Buckets {
		fmt.Println(aws.StringValue(bucket_item.Name), " : ", bucket_item.CreationDate)
		names[aws.StringValue(bucket_item.Name)]= struct{}{}
	}
	// private | public-read | public-read-write | authenticated-read
	// See https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL for details
	acl := "public-read"

	if _, ok := names[bucket]; !ok {
		_, err = svc.CreateBucket(&s3.CreateBucketInput{Bucket:&bucket})
		if err != nil {
			panic("create bucket error")
		}

		if err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{Bucket:&bucket}); err != nil {
			panic("failed to wait for bucket to exists")
		}


		//params := &s3.PutBucketAclInput{
		//	Bucket: &bucket,
		//	ACL: &acl,
		//}
		//
		//// Set bucket ACL
		//_, err := svc.PutBucketAcl(params)
		//if err != nil {
		//	panic(err)
		//}
		//
		//fmt.Println("Bucket " + bucket + " is now public")
	}

	rev_Item, err := svc.PutObject(&s3.PutObjectInput{
		Body: strings.NewReader("Hello World!"),
		Bucket: &bucket,
		Key: &keyfile,
		ACL: &acl,
	})

	if err != nil {
		panic("Failed to upload data to bucket")
	}

	fmt.Println(rev_Item)
	tag := rev_Item.ETag
	fmt.Println(tag)
	downloader := s3manager.NewDownloader(sess)
	newfile := "t1.txt"
	file, err := os.Create(newfile)
	numBytes, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:aws.String(keyfile),
	})

	if err != nil {
		panic("download error")
	}

	fmt.Println("Downlaod", file.Name(), numBytes, "bytes")

	fmt.Println("successfully upload")
}
