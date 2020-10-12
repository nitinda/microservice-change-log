package curd

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/utils/channels"
	"github.com/nitinda/microservice-change-log/logger"
)

func (r *respositoryUsersCRUD) ListAllUsers() ([]models.User, error) {
	var err error
	users := []models.User{}

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		// Model(&models.User{}).Limit(100).Find(&users).Error

		// Get all records with limit 100
		result := r.db.Limit(100).Find(&users)
		// SELECT * FROM "users" LIMIT 100

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
		return users, nil
	}
	return nil, err
}

func (r *respositoryUsersCRUD) ListUser(uid uint32) (models.User, error) {
	var err error
	user := models.User{}

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		// err = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&user).Error

		// if err != nil {
		// 	ch <- false
		// 	return
		// }

		// Get all matched records
		result := r.db.Where("id = ?", uid).Find(&user)

		// SELECT * FROM users WHERE id = 1;

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

	if gorm.IsRecordNotFoundError(err) {
		logger.Error.Println("User not found")
		return models.User{}, errors.New("User not found")
	}
	return models.User{}, err
}
