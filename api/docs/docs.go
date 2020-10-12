// Package docs changelog.
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
//     BasePath: /api/config
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
// swagger:meta
package docs

import (
	"github.com/nitinda/microservice-change-log/api/models"
)

// A list of ConfigLog
// swagger:response getConfiglogsResponse
type getConfiglogsResponseWrapper struct {
	// All current products
	// in: body
	Body []models.ConfigLog
}

// GenericError is a Generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Generic error message returned as a string
// swagger:response genericError
type genericErrorWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// Data structure representing a single changelog
// swagger:response configlogResponse
type configlogResponseWrapper struct {
	// Newly created Change-Log entry
	// in: body
	Body models.ConfigLog
}
