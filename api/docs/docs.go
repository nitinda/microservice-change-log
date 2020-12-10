// Package docs Changelog API
// Documentation
//
// the purpose of this application is to provide an application
// that will be used to record config chanegs applied to services.
//
// Terms Of Service:
//
// there are no TOS at this moment.
//
//	   Schemes: http
//     Host: localhost
//     BasePath: /api
//     Version: 1.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Nitin Das<nitinda@gmail.com>
//
//	   Consumes:
//     - application/json
//     - application/x-www-form-urlencoded
//
//     Produces:
//     - application/json
//
//     Security:
//     - OAuth2:
//     - Bearer:
//
//     SecurityDefinitions:
//     OAuth2:
//         type: oauth2
//         description: This API uses OAuth 2 with the Client Credntials grant flow.
//         tokenUrl: /oauth2/token
//         in: header
//         flow: application
//     Bearer:
//         type: apiKey
//         name: Authorization
//         in: header
//
//
// swagger:meta
package docs

import (
	"github.com/nitinda/microservice-change-log/api/models"
)

// GenericError is a Generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// Generic error message returned as a string
// swagger:response changelogErrorResponse
type changelogErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body struct {
		// Example: Token is expired or signature is invalid
		Message string `json:"unauthorized"`
	}
}

// Generic error message returned as a string
// swagger:response genericError
type genericErrorWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Data structure representing a single changelog entry
// swagger:parameters createChangeLog
type changelogRequestWrapper struct {
	// in:body
	Body struct {
		// the Service Team Name for the change log entry
		//
		// required: true
		// Unique: false
		// max length: 20
		// example: sales
		ServiceTeamName string `json:"ServiceTeamName"`

		// the Service Name for the change log entry
		//
		// required: true
		// Unique: false
		// max length: 20
		// example: tomcat
		ServiceName string `json:"ServiceName"`

		// the Username for the change log entry (Execution user)
		//
		// required: true
		// Unique: false
		// max length: 20
		// example: John Wick
		Username string `json:"Username"`

		// the Environment Name for the change log entry
		//
		// required: true
		// Unique: false
		// max length: 10
		// example: dev | test | prod
		EnvironmentName string `json:"EnvironmentName"`

		// the Commit Hash for the change log entry
		//
		// required: true
		// Unique: true
		// max length: 100
		// exmaple: 7b02a4c9b59sdfsda6eddffdsfsdf961b7d64bf32ebf
		CommitHash string `json:"CommitHash"`

		// the Release Info for the change log entry
		//
		// required: true
		// Unique: false
		// max length: 30
		// example: Release number 1.2.3
		ReleaseInfo string `json:"ReleaseInfo"`

		// the TypeOfChange for the change log entry
		//
		// required: true
		// Unique: false
		// max length: 10
		// example: config or release
		TypeOfChange string `gorm:"size:10;not null;type_of_change <> ''" json:"TypeOfChange"`

		// the Agent Info for the change log entry
		//
		// required: true
		// Unique: false
		// max length: 20
		// example: jenkins | gitlab
		AgentInfo string `json:"AgentInfo"`

		// the Message for the change log entry
		//
		// required: true
		// Unique: false
		// max length: 255
		// example: This is dummy release
		Message string `json:"Message"`
	}
}

// Data structure representing a single changelog
// swagger:response changelogResponse
type changelogResponseWrapper struct {
	// Newly created Change-Log entry
	// in: body
	Body models.ChangeLog
}

// Data structure representing a single teaminfo
// swagger:parameters createSessionToken
type createSessionTokenRquestWrapper struct {
	//
	// The Authorization to generate new session token
	// aBase64EncodeFunction(ApigeeKey:ApigeeSecret), Note the colon separating the two values.
	//
	// unique: true
	// required: true
	// in: header
	// example: Basic VG1iQ3NNaGZyVXJHCVKHNd3FZMzasHekZGUlRFT046TWMIZXCVNQAPk5xJSbQ==
	Authorization string `json:"Authorization"`

	// Content Type for request header
	//
	// required: true
	// in: header
	// default: application/x-www-form-urlencoded
	ContentType string `json:"Content-Type"`

	// the Authorization to generate new session token
	//
	// in:body
	Body struct {
		// The grant type for api authentication
		//
		// required: true
		// default: client_credentials
		GrantType string `json:"grant_type"`
	}
}

// Data structure representing create session token api response
// swagger:response createSessionTokenResponse
type createSessionTokenResponseWrapper struct {
	// Newly created Change-Log entry
	// in: body
	Body struct {
		// Example: Bearer
		TokenType string `json:"token_type"`

		// Example: Usf5LTdStGxER1GgHdGIj31VfPVmdNH
		AccessToken string `json:"access_token"`

		// Example: sdasdzoizcxhviuhIUIudsakfj
		CliendID string `json:"client_id"`

		// Example:
		Scope string `json:"scope"`

		// Example: 3599
		ExpiresIn int `json:"expires_in"`

		// Exmaple: 0
		RefeshCount int `json:"refresh_count"`

		// Exmple: 1607600275154
		IssuedAt int `json:"issued_at"`
	}
}

// A ValidationError is an error that is used when the required input fails validation.
// swagger:response createSessionTokenErrorResponse
type createSessionTokenErrorResponse struct {
	// To handle unauthorized access
	// in: body
	Body struct {
		FaultMessage struct {
			// Example: Invalid client identifier {0}
			FaultStringMessage string `json:"faultstring"`
			DetailMessage      struct {
				// Example: oauth.v2.InvalidClientIdentifier
				ErrorCodeMessage string `json:"errorcode"`
			} `json:"detail"`
		} `json:"fault"`
	}
}
