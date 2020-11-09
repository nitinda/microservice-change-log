package auto

import (
	"github.com/nitinda/microservice-change-log/api/models"
)

// teamInfo sample data
var teamInfo = []models.TeamInfo{
	models.TeamInfo{
		TeamName:     "sales",
		ClientSecret: "asdfasda87s6d876a8sd76a89d7669876zzxz8xc76z",
	}, models.TeamInfo{
		TeamName:     "inventory",
		ClientSecret: "sdasdzxcvzxcvzxc0897zxiuvpzx87c6v08p8127683",
	},
}

// changeLogs sample data
var changeLogs = []models.ChangeLog{
	models.ChangeLog{
		ServiceTeamName: "sales",
		ApplicationName: "hybris",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b5951ee6edddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
		Username:        "user1",
	}, models.ChangeLog{
		ServiceTeamName: "sales",
		ApplicationName: "hybris",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b5951ee6eddff23asdfdc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
		Username:        "user1",
	}, models.ChangeLog{
		ServiceTeamName: "sales",
		ApplicationName: "hybris",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b595asdf6eddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
		Username:        "user22",
	}, models.ChangeLog{
		ServiceTeamName: "sales",
		ApplicationName: "hybris",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9bqwe51ee6eddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
		Username:        "user3",
	}, models.ChangeLog{
		ServiceTeamName: "sales",
		ApplicationName: "hybris",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02aasdf9b5951ee6eddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
		Username:        "user11",
	}, models.ChangeLog{
		ServiceTeamName: "sales",
		ApplicationName: "atcom",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b0q1234c9b5951ee6eddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
		Username:        "user1",
	}, models.ChangeLog{
		ServiceTeamName: "sales",
		ApplicationName: "globalcache",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b5951ee6eddff23dcasdf7d64bf32ebf",
		AgentInfo:       "jenkins",
		Username:        "user4",
	}, models.ChangeLog{
		ServiceTeamName: "sales",
		ApplicationName: "atcom",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02asdf9b5951ee6eddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
		Username:        "user5",
	}, models.ChangeLog{
		ServiceTeamName: "sales",
		ApplicationName: "ngs",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b5951ee6eddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
		Username:        "user1",
	}, models.ChangeLog{
		ServiceTeamName: "inventoryData",
		ApplicationName: "inventoryData",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b5951ee6eddff23dc961basdf4bf32ebf",
		AgentInfo:       "jenkins",
		Username:        "user12",
	}, models.ChangeLog{
		ServiceTeamName: "inventory",
		ApplicationName: "inventoryData",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b5951ee6eddff23dc961basdf64bf32ebf",
		AgentInfo:       "jenkins",
		Username:        "agent11",
	}, models.ChangeLog{
		ServiceTeamName: "inventory",
		ApplicationName: "inventoryData",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b59dase6eddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
		Username:        "agent1",
	},
}
