package repositories

import "gin-mongodb-api/models"

type UserRepository interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	UpdateUser(*models.User) error
	GetAll() ([]*models.User, error)
	DeleteUser(*string) error
}
