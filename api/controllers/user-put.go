package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nitinda/microservice-change-log/api/database"
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/repository"
	"github.com/nitinda/microservice-change-log/api/repository/curd"
	"github.com/nitinda/microservice-change-log/api/responses"
)

// UpdateUser modify user information in the database
func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ValidateBody(rw, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
		return
	}

	db, er := database.DBConnectPostgres()
	dbSQL, ok := db.DB()
	if ok == nil {
		defer dbSQL.Close()
	}

	// defer db.Close()
	if er != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, er)
		return
	}

	repo := curd.NewRespositoryUsersCRUD(db)

	func(userRepository repository.UserReposiory) {
		rows, err := userRepository.UpdateUser(uint32(uid), user)
		if err != nil {
			responses.ValidateBody(rw, http.StatusBadGateway, err)
			return
		}
		responses.ToJSON(rw, http.StatusCreated, rows)
	}(repo)
}
