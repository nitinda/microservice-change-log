package curd

import (
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/utils/channels"
)

func (r *respositoryUsersCRUD) CreateNewUser(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.ValidateChannel(done) {
		return user, nil
	}
	return models.User{}, err
}
