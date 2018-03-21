package config

import (
	"fmt"
	"os/user"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// AWSParams saves config to to create an aws service clients
type AWSParams struct {
	Session *session.Session
}

// NewAWSParams creates a new AWSParams object
func NewAWSParams(awsProfile string) (*AWSParams, error) {
	var creds *credentials.Credentials

	if awsProfile != "" {
		currentUser, err := user.Current()
		if err != nil {
			return nil, err
		}

		creds = credentials.NewSharedCredentials(fmt.Sprintf("%s/.aws/credentials", currentUser.HomeDir), awsProfile)
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: creds,
	}))

	return &AWSParams{
		Session: sess,
	}, nil
}
