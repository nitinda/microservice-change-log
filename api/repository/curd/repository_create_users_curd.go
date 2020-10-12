package curd

import (
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/utils/channels"
	"github.com/nitinda/microservice-change-log/logger"
	"gorm.io/gorm"
)

func (r *respositoryUsersCRUD) CreateNewUser(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		// err = r.db.Model(&models.User{}).Create(&user).Error

		// if err != nil {
		// 	ch <- false
		// 	return
		// }

		// Update with conditions and model value
		result := r.db.Create(&user)
		// INSERT INTO "users" ("username","email","password","created_at","updated_at")
		// VALUES ('rowan','rowan@rowan.com','p12121212','2020-10-07 10:59:05.853','2020-10-07 10:59:05.853') RETURNING "id"

		if result.Error != nil {
			logger.Error.Println(result.Error, gorm.ErrRecordNotFound)
			ch <- false
			return
		}

		dbSQL, ok := result.DB()
		if ok == nil {
			defer dbSQL.Close()
		}

		ch <- true
	}(done)

	if channels.ValidateChannel(done) {
		return user, nil
	}
	return models.User{}, err
}
