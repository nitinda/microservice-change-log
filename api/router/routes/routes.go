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

	getR.HandleFunc("/api/healthcheck", middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(controllers.GetAPIStatus)))

	// swagger:route POST /api/changelog changelog createChangeLog
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
	//       Bearer:
	//
	//     Responses:
	//       200: changelogResponse
	//       401: changelogErrorResponse
	PostR.HandleFunc("/api/changelog", middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(controllers.CreateChangeLog)))

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
