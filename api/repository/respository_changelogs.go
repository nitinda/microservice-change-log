package repository

import (
	"github.com/nitinda/microservice-change-log/api/models"
)

type ChangeLogReposiory interface {
	CreateNewChangeLog(models.ChangeLog) (models.ChangeLog, error)
	// ListAllChangeLogs() ([]models.ChangeLog, error)
}
