package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nitinda/microservice-change-log/api/database"
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/repository"
	"github.com/nitinda/microservice-change-log/api/repository/curd"
	"github.com/nitinda/microservice-change-log/api/responses"
)

func CreateConfigLog(rw http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		return
	}

	configLog := models.ConfigLog{}
	err = json.Unmarshal(body, &configLog)
	if err != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		return
	}

	configLog.ConfigLogFieldCheck()

	err = configLog.ValidateConfigLog("")
	if err != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		return
	}

	db, er := database.DBConnect()
	defer db.Close()
	if er != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, er)
		return
	}

	repo := curd.NewRespositoryConfigLogsCRUD(db)

	func(configLogRepository repository.ConfigLogReposiory) {
		configLog, err := configLogRepository.CreateNewConfigLog(configLog)
		if err != nil {
			responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
			return
		}
		rw.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, configLog.ID))
		responses.ToJSON(rw, http.StatusCreated, configLog)
	}(repo)
}
