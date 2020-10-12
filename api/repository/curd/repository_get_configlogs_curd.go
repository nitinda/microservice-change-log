package curd

import (
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/utils/channels"
	"github.com/nitinda/microservice-change-log/logger"
	"gorm.io/gorm"
)

func (r *respositoryConfigLogsCRUD) ListAllConfigLogs() ([]models.ConfigLog, error) {
	var err error
	configLogs := []models.ConfigLog{}

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)

		result := r.db.Limit(100).Find(&configLogs)
		if result.Error != nil {
			logger.Error.Println(result.Error, gorm.ErrRecordNotFound)
			ch <- false
			return
		}

		dbSQL, ok := result.DB()
		if ok == nil {
			defer dbSQL.Close()
		}

		// err = r.db.Debug().Model(&models.ConfigLog{}).Limit(100).Find(&configLogs).Error

		// if err != nil {
		// 	ch <- false
		// 	return
		// }

		// if len(configLogs) > 0 {
		// 	for i, _ := range configLogs {
		// 		err = r.db.Debug().Model(&models.User{}).Where("id = ?", configLogs[i].UserID).Find(&configLogs[i].UserInfo).Error
		// 		if err != nil {
		// 			ch <- false
		// 			return
		// 		}
		// 	}
		// }

		ch <- true
	}(done)

	if channels.ValidateChannel(done) {
		return configLogs, nil
	}
	return nil, err
}

// func (r *respositoryConfigLogsCRUD) ListConfigLog(uid uint32) (models.ConfigLog, error) {
// 	var err error
// 	configLog := models.ConfigLog{}

// 	done := make(chan bool)
// 	go func(ch chan<- bool) {
// 		defer close(ch)
// 		err = r.db.Debug().Model(&models.ConfigLog{}).Where("id = ?", uid).Take(&configLog).Error

// 		if err != nil {
// 			ch <- false
// 			return
// 		}
// 		ch <- true
// 	}(done)

// 	if channels.ValidateChannel(done) {
// 		return configLog, nil
// 	}

// 	if gorm.IsRecordNotFoundError(err) {
// 		logger.Error.Println("Config Log entry not found")
// 		return models.ConfigLog{}, errors.New("Config Log entry not found")
// 	}
// 	return models.ConfigLog{}, err
// }
