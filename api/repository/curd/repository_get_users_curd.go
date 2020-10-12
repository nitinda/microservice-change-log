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
		err = r.db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error

		if err != nil {
			ch <- false
			return
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
		err = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&user).Error

		if err != nil {
			ch <- false
			return
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
