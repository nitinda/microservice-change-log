package routes

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gorilla/mux"
	"github.com/nitinda/microservice-change-log/api/controllers"
	"github.com/nitinda/microservice-change-log/api/middlewares"
)

// SetupRoutes set routes
func SetupRoutes(r *mux.Router) *mux.Router {

	// handlers for API
	getR := r.Methods(http.MethodGet).Subrouter()

	PostR := r.Methods(http.MethodPost).Subrouter()

	// swagger:route POST /api/token changelog createSessionToken
	// Create new session token entry
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Security:
	//       ApiKeyAuth: []
	//
	//     Responses:
	//       200: createSessionTokenResponse
	//       403: createSessionTokenErrorResponse
	PostR.HandleFunc("/api/token", middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(controllers.GenerateSessionToken)))

	// swagger:route POST /api/config changelog createChangeLog
	// Create new config log entry
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Security:
	//       BearerAuth: []
	//
	//     Responses:
	//       200: changelogResponse
	//       401: changelogErrorResponse
	PostR.HandleFunc("/api/config", middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(controllers.CreateChangeLog)))

	// DeleteR := r.Methods(http.MethodDelete).Subrouter()
	// DeleteR.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(controllers.DeleteUser)))

	// handler for documentation
	opts := middleware.RedocOpts{
		SpecURL: "/swagger.yaml",
		Path:    "/api/docs",
	}
	sh := middleware.Redoc(opts, nil)

	getR.Handle("/api/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./main/")))

	return r
}
