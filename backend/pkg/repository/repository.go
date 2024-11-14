package repository

import (
	"backend/pkg/model"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user model.User, role string) (int, error)
	GetUser(username, password string) ([]string, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Authorization: NewAuthPostgres(db)}
}
