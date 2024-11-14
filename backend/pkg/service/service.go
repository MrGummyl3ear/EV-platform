package service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User, role string) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) ([]string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
