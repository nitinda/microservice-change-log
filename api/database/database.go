package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nitinda/microservice-change-log/config"
)

// DBConnect will connect the database
func DBConnect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DB_DRIVER, config.DB_URL)

	if err != nil {
		return nil, err
	}
	return db, nil
}
