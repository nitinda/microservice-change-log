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

func CreateUser(rw http.ResponseWriter, r *http.Request) {

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

	user.UserFieldCheck()

	err = user.ValidateUser("")
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

	repo := curd.NewRespositoryUsersCRUD(db)

	func(userRepository repository.UserReposiory) {
		user, err := userRepository.CreateNewUser(user)
		if err != nil {
			responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
			return
		}
		rw.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, user.ID))
		responses.ToJSON(rw, http.StatusCreated, user)
	}(repo)

}
