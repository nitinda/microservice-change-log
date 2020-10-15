package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"gorm.io/gorm"

	"github.com/nitinda/microservice-change-log/api/security"
	"github.com/nitinda/microservice-change-log/logger"
)

// User struct abstract of user table in database
type User struct {
	ID         uint32      `gorm:"primary_key;auto_increment" json:"id"`
	Username   string      `gorm:"size:20;not null;unique" json:"username"`
	Email      string      `gorm:"size:50;not null;unique" json:"email"`
	Password   string      `gorm:"size:60;not null" json:"password"`
	CreatedAt  time.Time   `json:"created_on"`
	UpdatedAt  time.Time   `json:"updated_on"`
	ConfigLogs []ConfigLog `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"config_logs"`
}

// BeforeSave method encrypt the password
func (u *User) BeforeSave(db *gorm.DB) (err error) {
	hashedPassword, err := security.HashPassword(u.Password)
	if err != nil {
		log.Println(err)
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

// UserFieldCheck
func (u *User) UserFieldCheck() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) ValidateUser(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Username == "" {
			logger.Error.Println("Required Username")
			return errors.New("Required Username")
		}

		if u.Email == "" {
			logger.Error.Println("Required Email")
			return errors.New("Required Email")
		}

		err := checkmail.ValidateFormat(u.Email)
		if err != nil {
			logger.Error.Println("Invalid Email ID")
			return errors.New("Invalid Email")
		}
		return nil
	case "login":
		if u.Email == "" {
			logger.Error.Println("Required Email")
			return errors.New("Required Email")
		}

		err := checkmail.ValidateFormat(u.Email)
		if err != nil {
			logger.Error.Println("Invalid Email ID")
			return errors.New("Invalid Email")
		}

		if u.Password == "" {
			logger.Error.Println("Required Password")
			return errors.New("Required Password")
		}
		return nil
	default:
		if u.Username == "" {
			logger.Error.Println("Required Username")
			return errors.New("Required Username")
		}

		if u.Password == "" {
			logger.Error.Println("Required Password")
			return errors.New("Required Password")
		}

		if u.Email == "" {
			logger.Error.Println("Required Email")
			return errors.New("Required Email")
		}

		err := checkmail.ValidateFormat(u.Email)
		if err != nil {
			logger.Error.Println("Invalid Email ID")
			return errors.New("Invalid Email")
		}

		return nil
	}
}
