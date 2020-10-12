package curd

import (
	"time"

	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/utils/channels"
	"github.com/nitinda/microservice-change-log/logger"
	"gorm.io/gorm"
)

func (r *respositoryUsersCRUD) UpdateUser(uid uint32, user models.User) (int64, error) {

	// var err error
	var result *gorm.DB

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		// rs = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).UpdateColumns(
		// 	map[string]interface{}{
		// 		"username":   user.Username,
		// 		"email":      user.Email,
		// 		"updated_at": time.Now(),
		// 	},
		// )

		// if err != nil {
		// 	ch <- false
		// 	return
		// }

		// Update with conditions and model value
		result = r.db.Model(&models.User{}).Where("id = ?", uid).Updates(
			map[string]interface{}{
				"username":   user.Username,
				"email":      user.Email,
				"updated_at": time.Now(),
			},
		)
		// UPDATE users SET username='hello', updated_at='2013-11-17 21:34:10', email="asdfasdf" WHERE id=1 AND active=true;

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
		if result.Error != nil {
			return 0, result.Error
		}
		return result.RowsAffected, nil
	}

	return 0, result.Error
}
