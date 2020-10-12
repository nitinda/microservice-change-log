package curd

import (
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/utils/channels"
)

func (r *respositoryConfigLogsCRUD) CreateNewConfigLog(configLog models.ConfigLog) (models.ConfigLog, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.ConfigLog{}).Create(&configLog).Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.ValidateChannel(done) {
		return configLog, nil
	}
	return models.ConfigLog{}, err
}
