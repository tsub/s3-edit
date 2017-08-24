package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	myS3 "github.com/tsub/s3-edit/cli/s3"
)

// Edit directly a file on S3
func Edit(path myS3.Path) {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	bucket := aws.String(path.Bucket)
	key := aws.String(path.Key)

	getObjectInput := &s3.GetObjectInput {
		Bucket: bucket,
		Key: key,
	}
	res, err := svc.GetObject(getObjectInput)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(res.Body); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body := buf.Bytes()

	tempfilePath := "/tmp/" + *key
	if err := ioutil.WriteFile(tempfilePath, body, os.ModePerm); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer os.Remove(tempfilePath)

	cmd := exec.Command("sh", "-c", "nvim " + tempfilePath)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	changedFile, err := ioutil.ReadFile(tempfilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	putBody := string(changedFile[:])

	putObjectInput := &s3.PutObjectInput {
		Body: aws.ReadSeekCloser(strings.NewReader(putBody)),
		Bucket: bucket,
		Key: key,
	}
	if _, err := svc.PutObject(putObjectInput); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
