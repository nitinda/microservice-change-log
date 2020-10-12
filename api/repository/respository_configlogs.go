package repository

import (
	"github.com/nitinda/microservice-change-log/api/models"
)

type ConfigLogReposiory interface {
	CreateNewConfigLog(models.ConfigLog) (models.ConfigLog, error)
	ListAllConfigLogs() ([]models.ConfigLog, error)
	// ListConfigLog(uint32) (models.ConfigLog, error)
	// UpdateConfigLog(uint32, models.ConfigLog) (int64, error)
	// DeleteConfigLog(uint32) (int64, error)
}
