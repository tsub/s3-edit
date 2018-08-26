package s3

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// Object has object body and metadata
type Object struct {
	Body        []byte
	ContentType string
}

// GetObject download a file on S3
func GetObject(svc s3iface.S3API, path Path) Object {
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

	return Object{
		Body:        buf.Bytes(),
		ContentType: *res.ContentType,
	}
}

// PutObject upload a file to S3
func PutObject(svc s3iface.S3API, path Path, object Object) {
	input := &s3.PutObjectInput{
		ContentType: aws.String(object.ContentType),
		Body:        aws.ReadSeekCloser(bytes.NewReader(object.Body)),
		Bucket:      aws.String(path.Bucket),
		Key:         aws.String(path.Key),
	}
	if _, err := svc.PutObject(input); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
