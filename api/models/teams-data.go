package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/nitinda/microservice-change-log/api/security"
	"github.com/nitinda/microservice-change-log/logger"
	"gorm.io/gorm"
)

// TeamInfo struct abstract of Team data table in database
// swagger:model
type TeamInfo struct {
	TeamID uint32 `gorm:"primary_key;auto_increment" json:"TeamID"`

	// the Service Team Name to generate new session token
	//
	// required: true
	// Unique: true
	// max length: 20
	// example: sales
	TeamName string `gorm:"size:20;not null;unique" json:"TeamName"`

	// the Client Secret to generate new session token
	//
	// required: true
	// Unique: true
	// max length: 80
	// example: '$2a$10$m/22n7xmifpi1/rpsIzsIuY7.9walzEKloGCJF2ZpV.AElO83f2du'
	ClientSecret string `gorm:"size:80;not null;unique" json:"ClientSecret"`

	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

// BeforeSave method encrypt the secret
func (t *TeamInfo) BeforeSave(db *gorm.DB) (err error) {
	hashedSecret, err := security.HashSecret(t.ClientSecret)
	if err != nil {
		log.Println(err)
		return err
	}
	t.ClientSecret = string(hashedSecret)
	return nil
}

// TeamInfoFieldCheck method
func (t *TeamInfo) TeamInfoFieldCheck() {
	t.TeamID = 0
	t.TeamName = html.EscapeString(strings.TrimSpace(t.TeamName))
	t.ClientSecret = html.EscapeString(strings.TrimSpace(t.ClientSecret))
}

// ValidateTeamData method
func (t *TeamInfo) ValidateTeamData(action string) error {
	switch strings.ToLower(action) {
	case "token":
		if t.TeamName == "" {
			logger.Error.Println("Required Team Name")
			return errors.New("Required Team Name")
		}

		if t.ClientSecret == "" {
			logger.Error.Println("Required Secret Key")
			return errors.New("Required Secret Key")
		}

		if len(t.ClientSecret) <= 30 {
			logger.Error.Println("Required Secret Key length is less then 30 characters")
			return errors.New("Required Secret Key length is less then 30 characters")
		}
		return nil
	default:
		if t.TeamName == "" {
			logger.Error.Println("Required Team Name")
			return errors.New("Required Team Name")
		}

		if t.ClientSecret == "" {
			logger.Error.Println("Required Secret Key")
			return errors.New("Required Secret Key")
		}

		return nil
	}
}
