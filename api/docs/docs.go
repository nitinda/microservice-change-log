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
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Nitin Das<nitindas@gmail.com>
//
//	   Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - OAuth2: []
//
//     SecurityDefinitions:
//     OAuth2:
//         type: oauth2
//         description: This API uses OAuth 2 with the Client Credntials grant flow.
//         tokenUrl: /token
//         in: header
//         scopes:
//           email: userEmail
//           userinfo: username
//         flow: application
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

		// the Application Name for the change log entry
		//
		// required: true
		// Unique: false
		// max length: 20
		// example: ngs
		ApplicationName string `json:"ApplicationName"`

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
	// in:body
	Body struct {
		// the Service Team Name to generate new session token
		//
		// required: true
		// Unique: true
		// max length: 20
		// example: sales
		TeamName string `json:"TeamName"`

		// the Client Secret to generate new session token
		//
		// required: true
		// Unique: true
		// max length: 80
		// example: '$2a$10$m/22n7xmifpi1/rpsIzsIuY7.9walzEKloGCJF2ZpV.AElO83f2du'
		ClientSecret string `json:"ClientSecret"`
	}
}

// Data structure representing create session token api response
// swagger:response createSessionTokenResponse
type createSessionTokenResponseWrapper struct {
	// Newly created Change-Log entry
	// in: body
	Body struct {
		// Example: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDM3NTc2NzgsInRlYW1JRCI6MX0.uCRZnm6Fz6nI5_dV9lzRKejF3M5HB_MD5gRYuHG8sfVfPVmdNH
		Code string `json:"access_token"`
	}
}

// A ValidationError is an error that is used when the required input fails validation.
// swagger:response createSessionTokenErrorResponse
type createSessionTokenErrorResponse struct {
	// To handle unauthorized access
	// in: body
	Body struct {
		// Example: Access Denied
		Message string `json:"unauthorized"`
	}
}
