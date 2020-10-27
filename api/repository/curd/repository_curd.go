package curd

import "gorm.io/gorm"

// RespositoryChangeLogCRUD struct for getting and updating changelog in database
type RespositoryChangeLogCRUD struct {
	db *gorm.DB
}

// NewRespositoryChangeLogCRUD method
func NewRespositoryChangeLogCRUD(db *gorm.DB) *RespositoryChangeLogCRUD {
	return &RespositoryChangeLogCRUD{db}
}

// RespositoryTeamCRUD struct for getting and updating teaminfo in database
type RespositoryTeamCRUD struct {
	db *gorm.DB
}

// NewRespositoryTeamCRUD method returns a new RespositoryTeamCRUD handler
func NewRespositoryTeamCRUD(db *gorm.DB) *RespositoryTeamCRUD {
	return &RespositoryTeamCRUD{db}
}
