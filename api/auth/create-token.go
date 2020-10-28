package auth

import (
	"github.com/nitinda/microservice-change-log/config"
)

// CreateToken method
func CreateToken(teamName string, clientSecret string) (string, error) {

	serviceCognito := getAWSCognitoClient(config.AWS_REGION)

	// fmt.Println(serviceCognito)

	// initiateAuthRequest cognito
	tokenString, err := initiateAuthRequest(teamName, clientSecret, serviceCognito)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
