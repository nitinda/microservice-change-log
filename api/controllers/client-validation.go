package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nitinda/microservice-change-log/api/auth"
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/responses"
)

// GenerateSessionToken method
func GenerateSessionToken(rw http.ResponseWriter, r *http.Request) {

	var bodyBytes []byte
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		return
	}

	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	team := models.TeamInfo{}
	err = json.Unmarshal(bodyBytes, &team)
	if err != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		return
	}

	team.TeamInfoFieldCheck()

	err = team.ValidateTeamData("token")
	if err != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := auth.GenerateToken(team.TeamName, team.ClientSecret)
	if err != nil {
		//responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		responses.ToJSON(rw, http.StatusUnauthorized, map[string]string{"unauthorized": "Access Denied"})
		return
	}

	responses.ToJSON(rw, http.StatusOK, map[string]string{"access_token": token})
}
