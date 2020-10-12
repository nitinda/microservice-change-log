package auto

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nitinda/microservice-change-log/api/database"
	"github.com/nitinda/microservice-change-log/api/models"
	"github.com/nitinda/microservice-change-log/api/utils/console"
)

// LoadData will import data into database
func LoadData() {
	db, err := database.DBConnect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.ConfigLog{}, &models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.ConfigLog{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&models.ConfigLog{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal(err)
		}

		configLogs[i].UserID = users[i].ID
		err = db.Debug().Model(&models.ConfigLog{}).Create(&configLogs[i]).Error
		if err != nil {
			log.Fatal(err)
		}
		console.ToJSON(&configLogs[i])
	}

}
