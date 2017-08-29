package s3

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var sess = session.Must(session.NewSession())
var svc = s3.New(sess)

// GetObject download a file on S3
func GetObject(path Path) []byte {
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
func PutObject(path Path, body string) {
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
