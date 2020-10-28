package auto

import (
	"log"

	"github.com/nitinda/microservice-change-log/api/database"
	"github.com/nitinda/microservice-change-log/api/models"
)

// LoadData will import data into database
func LoadData() {
	db, err := database.DBConnectPostgres()
	if err != nil {
		log.Fatal(err)
	}

	dbSQL, ok := db.DB()
	if ok == nil {
		defer dbSQL.Close()
	}

	// Drop Table

	err = db.Migrator().DropTable(&models.ChangeLog{})
	if err != nil {
		log.Fatal(err)
	}

	// err = db.Migrator().DropTable(&models.TeamInfo{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Create Table

	err = db.AutoMigrate(&models.ChangeLog{})
	if err != nil {
		log.Fatal(err)
	}

	// err = db.AutoMigrate(&models.TeamInfo{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Insert data into table

	// err = db.Model(&models.TeamInfo{}).Create(&teamInfo).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = db.Model(&models.ChangeLog{}).Create(&changeLogs).Error
	if err != nil {
		log.Fatal(err)
	}
}
