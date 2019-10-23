package config

import (
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

// AWSParams saves config to to create an aws service clients
type AWSParams struct {
	Session *session.Session
}

// NewAWSParams creates a new AWSParams object
func NewAWSParams(awsProfile string) (*AWSParams, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:                 awsProfile,
		SharedConfigState:       session.SharedConfigEnable,
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	}))

	return &AWSParams{
		Session: sess,
	}, nil
}
