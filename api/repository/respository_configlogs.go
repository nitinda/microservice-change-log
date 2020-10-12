package repository

import (
	"github.com/nitinda/microservice-change-log/api/models"
)

type ConfigLogReposiory interface {
	CreateNewConfigLog(models.ConfigLog) (models.ConfigLog, error)
	ListAllConfigLogs() ([]models.ConfigLog, error)
}
