package curd

import (
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/utils/channels"
)

// CreateNewChangeLog method
func (r *RespositoryChangeLogCRUD) CreateNewChangeLog(changeLog models.ChangeLog) (models.ChangeLog, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Model(&models.ChangeLog{}).Create(&changeLog).Error

		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.ValidateChannel(done) {
		return changeLog, nil
	}
	return models.ChangeLog{}, err
}
