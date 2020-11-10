package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/nitinda/microservice-change-log/logger"
)

// ChangeLog defines the structure for an API config-log
// swagger:model
type ChangeLog struct {
	// the id for the change log entry - auto genrated
	//
	// required: false
	// Unique: true
	// Pattern: [0-9]+
	// min: 1
	ID uint64 `gorm:"primary_key;auto_increment" json:"ConfigEntryID"`

	// the Service Team Name for the change log entry
	//
	// required: true
	// Unique: false
	// max length: 20
	ServiceTeamName string `gorm:"size:20;not null" json:"ServiceTeamName"`

	// the Service Name for the change log entry
	//
	// required: true
	// Unique: false
	// max length: 20
	ServiceName string `gorm:"size:20;not null" json:"ServiceName"`

	// the Username for the change log entry (Execution user)
	//
	// required: true
	// Unique: false
	// max length: 20
	Username string `gorm:"size:20;not null;username <> ''" json:"Username"`

	// the Environment Name for the change log entry
	//
	// required: true
	// Unique: false
	// max length: 10
	EnvironmentName string `gorm:"size:20;not null;check:environment_name <> ''" json:"EnvironmentName"`

	// the Commit Hash for the change log entry
	//
	// required: true
	// Unique: true
	// max length: 100
	CommitHash string `gorm:"size:100;not null;unique;check:commit_hash <> ''" json:"CommitHash"`

	// the Release Info for the change log entry
	//
	// required: true
	// Unique: false
	// max length: 30
	ReleaseInfo string `gorm:"size:30;not null:check:release_info <> ''" json:"ReleaseInfo"`

	// the TypeOfChange for the change log entry
	//
	// required: true
	// Unique: false
	// max length: 20
	TypeOfChange string `gorm:"size:20;not null;check:type_of_change <> ''" json:"TypeOfChange"`

	// the Agent Info for the change log entry
	//
	// required: true
	// Unique: false
	// max length: 20
	AgentInfo string `gorm:"size:20;not null;check:agent_Info <> ''" json:"AgentInfo"`

	// the Message for the change log entry
	//
	// required: true
	// Unique: false
	// max length: 255
	Message string `gorm:"size:255;not null" json:"Message"`

	// the Creation Date for the change log entry - auto genrated
	//
	// required: false
	// Unique: true
	CreatedAt time.Time `json:"CreatedAt"`

	// the Modification Date for the change log entry - auto genrated
	// swagger:strfmt date-time
	//
	// required: false
	// Unique: true
	UpdatedAt time.Time `json:"UpdatedAt"`
}

// ChangeLogFieldCheck method
func (cl *ChangeLog) ChangeLogFieldCheck() {
	cl.ID = 0
	cl.ServiceTeamName = html.EscapeString(strings.TrimSpace(cl.ServiceTeamName))
	cl.ServiceName = html.EscapeString(strings.TrimSpace(cl.ServiceName))
	cl.Username = html.EscapeString(strings.TrimSpace(cl.Username))
	cl.CommitHash = html.EscapeString(strings.TrimSpace(cl.CommitHash))
	cl.AgentInfo = html.EscapeString(strings.TrimSpace(cl.AgentInfo))
	cl.ReleaseInfo = html.EscapeString(strings.TrimSpace(cl.ReleaseInfo))
	cl.TypeOfChange = html.EscapeString(strings.TrimSpace(cl.TypeOfChange))
	cl.Message = html.EscapeString(strings.TrimSpace(cl.Message))
	cl.CreatedAt = time.Now()
	cl.UpdatedAt = time.Now()
}

// ValidateChangeLog method
func (cl *ChangeLog) ValidateChangeLog(action string) error {
	if cl.ServiceTeamName == "" {
		logger.Error.Println("Required Service Team Name")
		return errors.New("Required Service Name")
	}

	if cl.ServiceName == "" {
		logger.Error.Println("Required Application Name")
		return errors.New("Required Team Name")
	}

	if cl.Message == "" {
		logger.Error.Println("Required Message")
		return errors.New("Required Message")
	}

	if cl.Username == "" {
		logger.Error.Println("Required Username")
		return errors.New("Required Username")
	}

	if cl.EnvironmentName == "" && cl.EnvironmentName == "dev|test|preprod|prod" {
		logger.Error.Println("Required Environment Name")
		return errors.New("Required Environment Name, Exmaple : dev or test or preprod or prod")
	}

	if cl.CommitHash == "" {
		logger.Error.Println("Required Commit Hash")
		return errors.New("Required Commit Hash")
	}

	if cl.ReleaseInfo == "" {
		logger.Error.Println("Required Release Information")
		return errors.New("Required Release Information")
	}

	if cl.TypeOfChange == "" {
		logger.Error.Println("Required TypeOfChange")
		return errors.New("Required TypeOfChange e.g - config or release")
	}

	if cl.AgentInfo == "" {
		logger.Error.Println("Required Agent Information")
		return errors.New("Required Agent Information e.g - jenkins or gitlab")
	}

	return nil
}
