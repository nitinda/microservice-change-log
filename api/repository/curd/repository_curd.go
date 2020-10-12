package curd

import "gorm.io/gorm"

type respositoryConfigLogsCRUD struct {
	db *gorm.DB
}

func NewRespositoryConfigLogsCRUD(db *gorm.DB) *respositoryConfigLogsCRUD {
	return &respositoryConfigLogsCRUD{db}
}

type respositoryUsersCRUD struct {
	db *gorm.DB
}

func NewRespositoryUsersCRUD(db *gorm.DB) *respositoryUsersCRUD {
	return &respositoryUsersCRUD{db}
}
