package auto

import (
	"github.com/nitinda/microservice-change-log/api/models"
)

var users = []models.User{
	models.User{
		Username: "hybris",
		Email:    "hybris@change.com",
		Password: "123453433",
	},
	// models.User{
	// 	Username: "Rowan",
	// 	Email:    "Rowan@change.com",
	// 	Password: "123453433",
	// },
}

// configLogs sample data
var configLogs = []models.ConfigLog{
	models.ConfigLog{
		Service: "hybris",
		Team:    "hybris-dev",
		Message: "This is test config entry",
		UserID:  1,
	}, models.ConfigLog{
		Service: "hybris",
		Team:    "hybris-dev",
		Message: "This is test config entry",
		UserID:  1,
	}, models.ConfigLog{
		Service: "hybris",
		Team:    "hybris-dev",
		Message: "This is test config entry",
		UserID:  1,
	}, models.ConfigLog{
		Service: "hybris",
		Team:    "hybris-dev",
		Message: "This is test config entry",
		UserID:  1,
	}, models.ConfigLog{
		Service: "hybris",
		Team:    "hybris-dev",
		Message: "This is test config entry",
		UserID:  1,
	}, models.ConfigLog{
		Service: "hybris",
		Team:    "hybris-dev",
		Message: "This is test config entry",
		UserID:  1,
	}, models.ConfigLog{
		Service: "hybris",
		Team:    "hybris-dev",
		Message: "This is test config entry",
		UserID:  1,
	}, models.ConfigLog{
		Service: "hybris",
		Team:    "hybris-dev",
		Message: "This is test config entry",
		UserID:  1,
	}, models.ConfigLog{
		Service: "hybris",
		Team:    "hybris-dev",
		Message: "This is test config entry",
		UserID:  1,
	}, models.ConfigLog{
		Service: "hybris",
		Team:    "hybris-dev",
		Message: "This is test config entry",
		UserID:  1,
	},
}
