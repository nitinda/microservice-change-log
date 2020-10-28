package curd

// import (
// 	"github.com/nitinda/microservice-change-log/api/models"
// 	"github.com/nitinda/microservice-change-log/api/utils/channels"
// 	"github.com/nitinda/microservice-change-log/logger"
// 	"gorm.io/gorm"
// )

// // ListAllChangeLogs method
// func (r *RespositoryChangeLogCRUD) ListAllChangeLogs() ([]models.ChangeLog, error) {
// 	var err error
// 	changeLogs := []models.ChangeLog{}

// 	done := make(chan bool)
// 	go func(ch chan<- bool) {
// 		defer close(ch)

// 		result := r.db.Limit(100).Find(&changeLogs)
// 		if result.Error != nil {
// 			logger.Error.Println(result.Error, gorm.ErrRecordNotFound)
// 			ch <- false
// 			return
// 		}

// 		dbSQL, ok := result.DB()
// 		if ok == nil {
// 			defer dbSQL.Close()
// 		}
// 		ch <- true
// 	}(done)

// 	if channels.ValidateChannel(done) {
// 		return changeLogs, nil
// 	}
// 	return nil, err
// }
