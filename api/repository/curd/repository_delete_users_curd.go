package curd

import (
	"github.com/jinzhu/gorm"
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/utils/channels"
)

func (r *respositoryUsersCRUD) DeleteUser(uid uint32) (int64, error) {
	var err error
	var rs *gorm.DB

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).Delete(&models.User{})

		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.ValidateChannel(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}

	return 0, rs.Error
}
