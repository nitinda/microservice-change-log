package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nitinda/microservice-change-log/api/database"
	"github.com/nitinda/microservice-change-log/api/repository"
	"github.com/nitinda/microservice-change-log/api/repository/curd"
	"github.com/nitinda/microservice-change-log/api/responses"
)

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ValidateBody(rw, http.StatusBadRequest, err)
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
		_, err := userRepository.DeleteUser(uint32(uid))
		if err != nil {
			responses.ValidateBody(rw, http.StatusBadGateway, err)
			return
		}
		rw.Header().Set("Entity", fmt.Sprintf("%d", uid))
		responses.ToJSON(rw, http.StatusNoContent, "")
	}(repo)
}
