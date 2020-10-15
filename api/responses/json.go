package responses

import (
	"encoding/json"
	"net/http"

	"github.com/nitinda/microservice-change-log/logger"
)

func ToJSON(rw http.ResponseWriter, statusCode int, data interface{}) {

	rw.Header().Set("Content-Type", "application/json")

	rw.WriteHeader(statusCode)

	if data != "" {
		err := json.NewEncoder(rw).Encode(data)
		if err != nil {
			logger.Info.Print(rw, "%s", err.Error())
		}
	}
}

func ValidateBody(rw http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		ToJSON(rw, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	ToJSON(rw, http.StatusBadRequest, nil)
}
