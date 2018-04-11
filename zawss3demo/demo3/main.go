package main

import (
	"net/http"
	"strings"
	"fmt"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
)
var (
	default_region = "ap-northeast-1"
	default_bucket = "beherrzbucket1"
	// private | public-read | public-read-write | authenticated-read
	// See https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL for details
	default_acl = "public-read"

	key_id = "AKIAIP65TRLVUOMNKALA"
	secret_key = "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"

	default_url = "https://s3-%s.amazonaws.com/%s/%s"

)

func main() {

	http.HandleFunc("/upload", UploadFile)

	http.ListenAndServe(":5051", nil)
}
func UploadFile(w http.ResponseWriter, r *http.Request) {
	defer uRecover(w)


//	var buf bytes.Buffer
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	file_name := header.Filename
	fmt.Println("file name: ", file_name)

	region := r.FormValue("region")
	if len(strings.Trim(region, " ")) == 0 {
		region = default_region
	}
	fmt.Println("regin: ", region)

	bucket := r.FormValue("bucket")
	if len(strings.Trim(bucket, " ")) == 0 {
		bucket = default_bucket
	}
	fmt.Println("bucket: ", bucket)


//	io.Copy(&buf, file)

	svc := InitS3(region, bucket)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Body: file,
		Bucket: aws.String(bucket),
		Key: aws.String(file_name),
		ACL: aws.String(default_acl),
	})

	if err != nil {
		panic("Failed to upload data to bucket")
	}

	rev_url := fmt.Sprintf(default_url, region, bucket, file_name )

	ResponseJson(w, &UploadRevData{Url: rev_url})

}

func uRecover(w http.ResponseWriter) {
	if err := recover(); err != nil {
		fmt.Println(err)
		ResponseJson(w, &ErrorInfo{ErrorCode:"505", ErrorInfo:"server error"})
	}
}

type ErrorInfo struct {
	ErrorCode string
	ErrorInfo string
}

func InitS3(region, bucket string) *s3.S3{
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(key_id, secret_key, ""),

	})
	if err != nil {
		panic(err)
	}

	svc := s3.New(sess)
	_, err = svc.HeadBucket(&s3.HeadBucketInput{Bucket: &bucket})
	if err != nil {
		_, err = svc.CreateBucket(&s3.CreateBucketInput{Bucket:&bucket})
		if err != nil {
			panic("create bucket error")
		}

		if err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{Bucket:&bucket}); err != nil {
			panic("failed to wait for bucket to exists")
		}
	}

	return svc
}

type UploadRevData struct {
	Url string
}


func ResponseJson(w http.ResponseWriter, data interface{}) {
	rdata, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rdata)
}