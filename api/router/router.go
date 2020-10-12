package router

import (
	"github.com/gorilla/mux"
	"github.com/nitinda/microservice-change-log/api/router/routes"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}
