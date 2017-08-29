package s3

import (
	"fmt"
	"net/url"
	"os"
)

// A Path have S3 bucket and key
type Path struct {
	Bucket string
	Key    string
}

// ParsePath parse a file path of S3, and return Path struct
func ParsePath(path string) Path {
	parsedURL, err := url.ParseRequestURI(path)
	if err != nil {
		fmt.Println("invalid S3 path")
		os.Exit(1)
	}

	if isInvalidS3Path(*parsedURL) {
		fmt.Println("invalid S3 path")
		os.Exit(1)
	}

	return Path{
		Bucket: parsedURL.Host,
		Key:    parsedURL.Path,
	}
}

func isInvalidS3Path(url url.URL) bool {
	return url.Host == "" || url.Path == "" || url.Path == "/"
}
