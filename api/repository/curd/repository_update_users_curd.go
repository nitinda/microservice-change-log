package curd

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/utils/channels"
)

func (r *respositoryUsersCRUD) UpdateUser(uid uint32, user models.User) (int64, error) {

	var err error
	var rs *gorm.DB

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).UpdateColumns(
			map[string]interface{}{
				"username":   user.Username,
				"email":      user.Email,
				"updated_at": time.Now(),
			},
		)

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
