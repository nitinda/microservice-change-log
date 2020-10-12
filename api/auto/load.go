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

	err = db.Migrator().DropTable(&models.ConfigLog{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrator().DropTable(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	// defer db.Close()

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.ConfigLog{})
	if err != nil {
		log.Fatal(err)
	}

	// Insert data
	err = db.Model(&models.User{}).Create(&users).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Model(&models.ConfigLog{}).Create(&configLogs).Error
	if err != nil {
		log.Fatal(err)
	}
}
