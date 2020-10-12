package auto

import (
	"github.com/nitinda/microservice-change-log/api/models"
)

var users = []models.User{
	models.User{
		Username: "Neville",
		Email:    "Neville@change.com",
		Password: "123453433",
	},
	// models.User{
	// 	Username: "Rowan",
	// 	Email:    "Rowan@change.com",
	// 	Password: "123453433",
	// },
}

var configLogs = []models.ConfigLog{
	models.ConfigLog{
		Service: "hybris",
		Team:    "nr-cpt",
		Message: "This is test config entry",
	},
}
