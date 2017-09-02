package s3

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// GetObject download a file on S3
func GetObject(svc s3iface.S3API, path Path) []byte {
	input := &s3.GetObjectInput{
		Bucket: aws.String(path.Bucket),
		Key:    aws.String(path.Key),
	}
	res, err := svc.GetObject(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(res.Body); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return buf.Bytes()
}

// PutObject upload a file to S3
func PutObject(svc s3iface.S3API, path Path, body string) {
	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(strings.NewReader(body)),
		Bucket: aws.String(path.Bucket),
		Key:    aws.String(path.Key),
	}
	if _, err := svc.PutObject(input); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
