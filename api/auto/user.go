package auto

import (
	"log"

	"github.com/nitinda/microservice-change-log/api/database"
	"github.com/nitinda/microservice-change-log/logger"
)

// CreateUser will create user for grafana
func CreateUser() {
	db, err := database.DBConnectPostgres()
	if err != nil {
		log.Fatal(err)
	}

	dbSQL, ok := db.DB()
	if ok == nil {
		defer dbSQL.Close()
	}

	// 	result := db.First(&user)
	// result.RowsAffected // returns found records count
	// result.Error        // returns error

	// // check error ErrRecordNotFound
	// errors.Is(result.Error, gorm.ErrRecordNotFound)

	result := db.Exec("SELECT 1 FROM pg_roles WHERE rolname='grafana'")
	if result.RowsAffected == 0 {

		err = db.Exec("CREATE USER grafana WITH PASSWORD 'grafana';").Error
		if err != nil {
			logger.Error.Println(err)
		}

		err = db.Exec("GRANT CONNECT ON DATABASE postgres TO grafana;").Error
		if err != nil {
			logger.Error.Println(err)
		}

		err = db.Exec("GRANT USAGE ON SCHEMA public TO grafana").Error
		if err != nil {
			logger.Error.Println(err)
		}

		err = db.Exec("GRANT SELECT ON change_logs TO grafana").Error
		if err != nil {
			logger.Error.Println(err)
		}
	}

	// err = db.Exec("GRANT CONNECT ON DATABASE postgres TO grafana;").Error
	// if err != nil {
	// 	logger.Error.Println(err)
	// }
}
