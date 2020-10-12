package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/nitinda/microservice-change-log/logger"
)

// ConfigLog defines the structure for an API config-log
// swagger:model
type ConfigLog struct {
	// the id for the config log entry - auto genrated
	//
	// required: false
	// Unique: true
	// Pattern: [0-9]+
	// min: 1
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`

	// the service name for the config log entry
	//
	// required: true
	// Unique: false
	// max length: 20
	Service string `gorm:"size:20;not null" json:"service"`

	// the team name for the config log entry
	//
	// required: true
	// Unique: false
	// max length: 20
	Team string `gorm:"size:20;not null" json:"team"`

	UserInfo []User `json:"user"`
	UserID   uint32 `gorm:"not null" json:"user_id"`

	// the message for the config log entry
	//
	// required: true
	// max length: 255
	Message string `gorm:"size:255;not null" json:"message"`

	// the creation date for the config log entry - auto genrated
	//
	// required: false
	// Unique: true
	CreatedAt time.Time `json:"created_on"`

	// the modification date for the config log entry - auto genrated
	//
	// required: false
	// Unique: true
	UpdatedAt time.Time `json:"updated_on"`
}

// UserFieldCheck
func (cl *ConfigLog) ConfigLogFieldCheck() {
	cl.ID = 0
	cl.Service = html.EscapeString(strings.TrimSpace(cl.Service))
	cl.Team = html.EscapeString(strings.TrimSpace(cl.Team))
	cl.Message = html.EscapeString(strings.TrimSpace(cl.Message))
	cl.CreatedAt = time.Now()
	cl.UpdatedAt = time.Now()
}

func (cl *ConfigLog) ValidateConfigLog(action string) error {
	if cl.Service == "" {
		logger.Error.Println("Required Service Name")
		return errors.New("Required Service Name")
	}

	if cl.Team == "" {
		logger.Error.Println("Required Team Name")
		return errors.New("Required Team Name")
	}

	if cl.Message == "" {
		logger.Error.Println("Required Message")
		return errors.New("Required Message")
	}

	return nil
}
