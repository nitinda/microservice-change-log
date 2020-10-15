package auth

import (
	"github.com/nitinda/microservice-change-log/api/database"
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/security"
	"github.com/nitinda/microservice-change-log/api/utils/channels"
	"github.com/nitinda/microservice-change-log/logger"
	"gorm.io/gorm"
)

func SignIn(email, password string) (string, error) {

	user := models.User{}
	var err error
	var db *gorm.DB

	done := make(chan bool)

	go func(ch chan<- bool) {

		db, err = database.DBConnectPostgres()
		if err != nil {
			ch <- false
			return
		}

		// Get all matched records
		result := db.Where("email = ?", email).Find(&user)

		// SELECT * FROM users WHERE email = 'test@123';

		if result.Error != nil {
			logger.Error.Println(result.Error, gorm.ErrRecordNotFound)
			ch <- false
			return
		}

		dbSQL, ok := result.DB()
		if ok == nil {
			defer dbSQL.Close()
		}

		err = security.VarrifyPassword(user.Password, password)
		if err != nil {
			ch <- false
			return
		}

		ch <- true

	}(done)

	if channels.ValidateChannel(done) {
		return CreateToken(user.ID)
	}

	return "", err

}
