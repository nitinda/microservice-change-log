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
		ServiceTeamName: "hybris",
		ApplicationName: "sale",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b5951ee6edddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
	}, models.ChangeLog{
		ServiceTeamName: "hybris",
		ApplicationName: "sale",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b5951ee6eddff23asdfdc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
	}, models.ChangeLog{
		ServiceTeamName: "hybris",
		ApplicationName: "sale",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b595asdf6eddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
	}, models.ChangeLog{
		ServiceTeamName: "hybris",
		ApplicationName: "sale",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9bqwe51ee6eddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
	}, models.ChangeLog{
		ServiceTeamName: "hybris",
		ApplicationName: "sale",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02aasdf9b5951ee6eddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
	}, models.ChangeLog{
		ServiceTeamName: "hybris",
		ApplicationName: "sale",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b0q1234c9b5951ee6eddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
	}, models.ChangeLog{
		ServiceTeamName: "hybris",
		ApplicationName: "sale",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b5951ee6eddff23dcasdf7d64bf32ebf",
		AgentInfo:       "jenkins",
	}, models.ChangeLog{
		ServiceTeamName: "hybris",
		ApplicationName: "sale",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02asdf9b5951ee6eddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
	}, models.ChangeLog{
		ServiceTeamName: "hybris",
		ApplicationName: "sale",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b5951ee6eddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
	}, models.ChangeLog{
		ServiceTeamName: "inventoryData",
		ApplicationName: "inventory-dev",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b5951ee6eddff23dc961basdf4bf32ebf",
		AgentInfo:       "jenkins",
	}, models.ChangeLog{
		ServiceTeamName: "inventory",
		ApplicationName: "inventory-dev",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b5951ee6eddff23dc961basdf64bf32ebf",
		AgentInfo:       "jenkins",
	}, models.ChangeLog{
		ServiceTeamName: "inventory",
		ApplicationName: "inventory-dev",
		Message:         "This is test config entry",
		EnvironmentName: "prod",
		ReleaseInfo:     "1.2.4",
		CommitHash:      "7b02a4c9b59dase6eddff23dc961b7d64bf32ebf",
		AgentInfo:       "jenkins",
	},
}
