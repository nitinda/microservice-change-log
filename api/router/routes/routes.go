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
	getR.HandleFunc("/api/users", middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(controllers.GetUsers)))

	// swagger:route GET /api/config changelog listChangeLog
	// Return a list of config logs from the database
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//       default: genericError
	//       200: getConfiglogsResponse
	getR.HandleFunc("/api/config", middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(controllers.GetConfigLogs)))

	getR.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(controllers.GetUser)))

	PutR := r.Methods(http.MethodPut).Subrouter()
	PutR.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(controllers.UpdateUser)))

	PostR := r.Methods(http.MethodPost).Subrouter()
	PostR.HandleFunc("/api/users", middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(controllers.CreateUser)))

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
	//     Responses:
	//       200: configlogResponse
	//       422: errorValidation
	//       501: errorResponse
	PostR.HandleFunc("/api/config", middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(controllers.CreateConfigLog))))

	// User Login
	PostR.HandleFunc("/api/login", middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(controllers.Login))))

	DeleteR := r.Methods(http.MethodDelete).Subrouter()
	DeleteR.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(controllers.DeleteUser)))

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
