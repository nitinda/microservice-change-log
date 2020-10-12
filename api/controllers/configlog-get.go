package controllers

import (
	"net/http"

	"github.com/nitinda/microservice-change-log/api/database"
	"github.com/nitinda/microservice-change-log/api/repository"
	"github.com/nitinda/microservice-change-log/api/repository/curd"
	"github.com/nitinda/microservice-change-log/api/responses"
)

// GetConfigLogs handles GET requests and returns all config logs from
func GetConfigLogs(rw http.ResponseWriter, r *http.Request) {
	db, er := database.DBConnect()
	defer db.Close()
	if er != nil {
		responses.ValidateBody(rw, http.StatusUnprocessableEntity, er)
		return
	}

	repo := curd.NewRespositoryConfigLogsCRUD(db)

	func(configLogRepository repository.ConfigLogReposiory) {
		configLogs, err := configLogRepository.ListAllConfigLogs()
		if err != nil {
			responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
			return
		}
		responses.ToJSON(rw, http.StatusCreated, configLogs)
	}(repo)
}

// func GetConfigLog(rw http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	uid, err := strconv.ParseUint(vars["id"], 10, 32)
// 	if err != nil {
// 		responses.ValidateBody(rw, http.StatusBadRequest, err)
// 		return
// 	}

// 	db, er := database.DBConnect()
// 	defer db.Close()
// 	if er != nil {
// 		responses.ValidateBody(rw, http.StatusUnprocessableEntity, er)
// 		return
// 	}

// 	repo := curd.NewRespositoryConfigLogsCRUD(db)

// 	func(configLogRepository repository.ConfigLogReposiory) {
// 		configLog, err := configLogRepository.ListConfigLog(uint32(uid))
// 		if err != nil {
// 			responses.ValidateBody(rw, http.StatusBadGateway, err)
// 			return
// 		}
// 		responses.ToJSON(rw, http.StatusCreated, configLog)
// 	}(repo)
// }
