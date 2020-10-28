package controllers

// import (
// 	"net/http"

// 	"github.com/nitinda/microservice-change-log/api/database"
// 	"github.com/nitinda/microservice-change-log/api/repository"
// 	"github.com/nitinda/microservice-change-log/api/repository/curd"
// 	"github.com/nitinda/microservice-change-log/api/responses"
// )

// // GetChangeLogs handles GET requests and returns all change logs from
// func GetChangeLogs(rw http.ResponseWriter, r *http.Request) {
// 	db, er := database.DBConnectPostgres()

// 	dbSQL, ok := db.DB()
// 	if ok == nil {
// 		defer dbSQL.Close()
// 	}
// 	// defer db.Close()
// 	if er != nil {
// 		responses.ValidateBody(rw, http.StatusUnprocessableEntity, er)
// 		return
// 	}

// 	repo := curd.NewRespositoryChangeLogCRUD(db)

// 	func(changeLogRepository repository.ChangeLogReposiory) {
// 		changeLogs, err := changeLogRepository.ListAllChangeLogs()
// 		if err != nil {
// 			responses.ValidateBody(rw, http.StatusUnprocessableEntity, err)
// 			return
// 		}
// 		responses.ToJSON(rw, http.StatusCreated, changeLogs)
// 	}(repo)
// }
