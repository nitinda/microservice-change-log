package curd

import "gorm.io/gorm"

type respositoryConfigLogsCRUD struct {
	db *gorm.DB
}

// RespositoryChangeLogCRUD struct for getting and updating changelog in database
type RespositoryChangeLogCRUD struct {
	db *gorm.DB
}

// NewRespositoryChangeLogCRUD method
func NewRespositoryChangeLogCRUD(db *gorm.DB) *RespositoryChangeLogCRUD {
	return &RespositoryChangeLogCRUD{db}
}
