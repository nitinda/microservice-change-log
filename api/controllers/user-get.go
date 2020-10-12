package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nitinda/microservice-change-log/api/database"
	"github.com/nitinda/microservice-change-log/api/repository"
	"github.com/nitinda/microservice-change-log/api/repository/curd"
	"github.com/nitinda/microservice-change-log/api/responses"
)

// GetUsers list all user from database
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	db, er := database.DBConnectPostgres()
	// defer db.Close()
	if er != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, er)
		return
	}

	repo := curd.NewRespositoryUsersCRUD(db)

	func(userRepository repository.UserReposiory) {
		users, err := userRepository.ListAllUsers()
		if err != nil {
			responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
			return
		}
		responses.ToJSON(rw, http.StatusCreated, users)
	}(repo)
}

// GetUser list one user from database
func GetUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ValidateBody(rw, http.StatusBadRequest, err)
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
		user, err := userRepository.ListUser(uint32(uid))
		if err != nil {
			responses.ValidateBody(rw, http.StatusBadGateway, err)
			return
		}
		responses.ToJSON(rw, http.StatusCreated, user)
	}(repo)

}
