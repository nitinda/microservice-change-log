package auth

import (
	"github.com/nitinda/microservice-change-log/api/database"
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/utils/channels"
	"github.com/nitinda/microservice-change-log/logger"
	"gorm.io/gorm"
)

// GenerateToken method
func GenerateToken(teamName, teamSecret string) (string, error) {
	team := models.TeamInfo{}
	var err error
	var db *gorm.DB

	done := make(chan bool)

	go func(ch chan<- bool) {

		db, err = database.DBConnectPostgres()
		if err != nil {
			ch <- false
			return
		}

		// Get matched records for provided team and
		result := db.Where("team_name = ? AND client_secret <> ?", teamName, teamSecret).Find(&team)

		// SELECT * FROM team_info WHERE team_name = 'test@123';

		if result.Error != nil || result.RowsAffected == 0 {
			logger.Error.Println(result.Error, gorm.ErrRecordNotFound)
			err = gorm.ErrRecordNotFound
			ch <- false
			return
		}

		// Database Close
		dbSQL, ok := result.DB()
		if ok == nil {
			defer dbSQL.Close()
		}

		ch <- true
	}(done)

	if channels.ValidateChannel(done) {
		return CreateToken(team.TeamName, team.ClientSecret)
	}

	return "", err
}
