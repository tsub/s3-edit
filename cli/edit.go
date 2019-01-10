package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3"
	myS3 "github.com/tsub/s3-edit/cli/s3"
	"github.com/tsub/s3-edit/config"
)

// Edit directly a file on S3
func Edit(path myS3.Path, params *config.AWSParams) {
	svc := s3.New(params.Session)

	object := myS3.GetObject(svc, path)

	tempDirPath, tempfilePath := createTempfile(path, object.Body)
	defer os.RemoveAll(tempDirPath)

	editedBody := editFile(tempfilePath)
	object.Body = []byte(editedBody)
	myS3.PutObject(svc, path, object)
}

func createTempfile(path myS3.Path, body []byte) (tempDirPath string, tempfilePath string) {
	tempDirPath, err := ioutil.TempDir("/tmp", "s3-edit")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	keys := strings.Split(path.Key, "/")
	fileName := keys[len(keys)-1]
	tempfilePath = tempDirPath + "/" + fileName

	if err := ioutil.WriteFile(tempfilePath, body, os.ModePerm); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}

func editFile(path string) string {
	command := getDefaultEditor() + " " + path

	cmd := exec.Command("sh", "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	changedFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(changedFile[:])
}

func getDefaultEditor() string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return "vi"
	}
	return editor
}
