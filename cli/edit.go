package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	myS3 "github.com/tsub/s3-edit/cli/s3"
)

// Edit directly a file on S3
func Edit(path myS3.Path) {
	body := myS3.GetObject(path)

	tempfilePath := createTempfile(path, body)
	defer os.Remove(tempfilePath)

	editedBody := editFile(tempfilePath)
	myS3.PutObject(path, editedBody)
}

func createTempfile(path myS3.Path, body []byte) (tempfilePath string) {
	keys := strings.Split(path.Key, "/")
	fileName := keys[len(keys)-1]
	tempfilePath = "/tmp/" + fileName

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
	return os.Getenv("EDITOR")
}
