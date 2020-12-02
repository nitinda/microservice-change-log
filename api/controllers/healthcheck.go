package controllers

import (
	"net/http"

	"github.com/nitinda/microservice-change-log/api/responses"
)

// GetAPIStatus method
func GetAPIStatus(rw http.ResponseWriter, r *http.Request) {

	responses.ToJSON(rw, http.StatusOK, map[string]string{"Service": "Ok"})

}
