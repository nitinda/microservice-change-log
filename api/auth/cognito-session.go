package auth

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// getAWSCognitoClient method
func getAWSCognitoClient(region string) *cognito.Client {

	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	// Set the AWS Region that the service clients should use
	cfg.Region = region

	// Using the Config value, create the Cognito client
	svc := cognito.New(cfg)

	return svc
}
