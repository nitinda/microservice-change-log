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

// CreateChangeLog method
func CreateChangeLog(rw http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		return
	}

	changeLog := models.ChangeLog{}
	err = json.Unmarshal(body, &changeLog)
	if err != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		return
	}

	changeLog.ChangeLogFieldCheck()

	err = changeLog.ValidateChangeLog("")
	if err != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		return
	}

	db, er := database.DBConnectPostgres()

	dbSQL, ok := db.DB()
	if ok == nil {
		defer dbSQL.Close()
	}

	if er != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, er)
		return
	}

	repo := curd.NewRespositoryChangeLogCRUD(db)

	func(changeLogRepository repository.ChangeLogReposiory) {
		changeLog, err := changeLogRepository.CreateNewChangeLog(changeLog)
		if err != nil {
			responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
			return
		}
		rw.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, changeLog.ID))
		responses.ToJSON(rw, http.StatusCreated, changeLog)
	}(repo)
}
