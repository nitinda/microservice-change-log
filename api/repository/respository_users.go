package repository

import (
	"github.com/nitinda/microservice-change-log/api/models"
)

type UserReposiory interface {
	CreateNewUser(models.User) (models.User, error)
	ListAllUsers() ([]models.User, error)
	ListUser(uint32) (models.User, error)
	UpdateUser(uint32, models.User) (int64, error)
	DeleteUser(uint32) (int64, error)
}
