package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nitinda/microservice-change-log/api/auth"
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/responses"
)

// GenerateSessionToken method
func GenerateSessionToken(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		return
	}

	team := models.TeamInfo{}
	err = json.Unmarshal(body, &team)
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
