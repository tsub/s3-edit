package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

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
	tempfilePath = "/tmp/" + path.Key

	if err := ioutil.WriteFile(tempfilePath, body, os.ModePerm); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}

func editFile(path string) string {
	cmd := exec.Command("sh", "-c", "nvim " + path)
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
