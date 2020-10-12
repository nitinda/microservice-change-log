package database

import "gorm.io/gorm"

func DBClose(db *gorm.DB) {
	dbSQL, ok := db.DB()
	if ok != nil {
		defer dbSQL.Close()
	}
}
