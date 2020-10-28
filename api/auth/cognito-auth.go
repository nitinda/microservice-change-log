package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	cognitosrp "github.com/alexrudd/cognito-srp/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/nitinda/microservice-change-log/config"
	"github.com/nitinda/microservice-change-log/logger"
)

// Special thanks to - https://github.com/mura123yasu/go-cognito

// initiateAuthRequest method
func initiateAuthRequest(teamName, clientSecret string, client *cognito.Client) (string, error) {

	var tokenString string

	cognitoClientSecret := config.COGNITO_CLIENT_SECRET
	csrp, _ := cognitosrp.NewCognitoSRP(teamName, clientSecret, config.COGNITO_USER_POOL_ID, config.COGNITO_CLIENT_ID, &cognitoClientSecret)

	// initiate auth
	initiateAuthInputParams := &cognito.InitiateAuthInput{
		AuthFlow:       cognito.AuthFlowTypeUserSrpAuth,
		ClientId:       aws.String(csrp.GetClientId()),
		AuthParameters: csrp.GetAuthParams(),
	}
	req := client.InitiateAuthRequest(initiateAuthInputParams)

	resp, reqError := req.Send(context.TODO())
	if reqError != nil {
		logger.Error.Println(reqError)
		return "", errors.New("InitiateAuthRequest failed with ID Provider")
	}

	// respond to password verifier challenge
	if resp.ChallengeName == cognito.ChallengeNameTypePasswordVerifier {
		challengeResponses, challengeResoncesError := csrp.PasswordVerifierChallenge(resp.ChallengeParameters, time.Now())
		if challengeResoncesError != nil {
			log.Fatal(challengeResoncesError)
		}

		respondToAuthChallengeInputParams := &cognito.RespondToAuthChallengeInput{
			ChallengeName:      cognito.ChallengeNameTypePasswordVerifier,
			ChallengeResponses: challengeResponses,
			ClientId:           aws.String(csrp.GetClientId()),
		}

		request := client.RespondToAuthChallengeRequest(respondToAuthChallengeInputParams)

		resp, err := request.Send(context.TODO())
		if err != nil {
			// Casting to the awserr.Error type will allow you to inspect the error
			// code returned by the service in code. The error code can be used
			// to switch on context specific functionality. In this case a context
			// specific error message is printed to the user based on the bucket
			// and key existing.

			var errString string

			if awsErr, ok := err.(awserr.Error); ok {
				switch awsErr.Code() {
				case cognito.ErrCodeNotAuthorizedException:
					logger.Error.Println(cognito.ErrCodeNotAuthorizedException, ": ", awsErr.Message())
					errString = fmt.Sprintf("%v%v%v", cognito.ErrCodeNotAuthorizedException, ": ", awsErr.Message())
				}
			}
			return "", errors.New(errString)
		}

		// extract the AccessToken
		tokenString = *resp.AuthenticationResult.AccessToken
	}

	// fmt.Println(resp, reqError)
	return tokenString, nil
}
