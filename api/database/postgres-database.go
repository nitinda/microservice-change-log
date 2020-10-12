package database

import (
	"github.com/nitinda/microservice-change-log/config"
	"github.com/nitinda/microservice-change-log/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBConnectPostgres to connect the database
func DBConnectPostgres() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DB_URL), &gorm.Config{
		Logger:                                   logger.DBLoggerInfo,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: false,
		PrepareStmt:                              false,
	})

	if err != nil {
		return nil, err
	}
	return db, nil
}
