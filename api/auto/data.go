package auto

import (
	"github.com/nitinda/microservice-change-log/api/models"
)

var users = []models.User{
	models.User{
		Username: "sale",
		Email:    "sale@changeapi.com",
		Password: "123453433",
	},
	models.User{
		Username: "inventory",
		Email:    "inventory@changeapi.com",
		Password: "97878789778",
	},
	models.User{
		Username: "postman",
		Email:    "postman@changeapi.com",
		Password: "4545454544",
	},
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
	},
}
