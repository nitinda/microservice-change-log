package curd

import "github.com/jinzhu/gorm"

type respositoryUsersCRUD struct {
	db *gorm.DB
}

type respositoryConfigLogsCRUD struct {
	db *gorm.DB
}

func NewRespositoryUsersCRUD(db *gorm.DB) *respositoryUsersCRUD {
	return &respositoryUsersCRUD{db}
}

func NewRespositoryConfigLogsCRUD(db *gorm.DB) *respositoryConfigLogsCRUD {
	return &respositoryConfigLogsCRUD{db}
}
