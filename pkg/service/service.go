package service

import (
	model "github.com/vnSasa/music-market-api/model"
	"github.com/vnSasa/music-market-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) error
	GenerateToken(login, password string) (*model.TokenDetails, error)
	ParseToken(accessToken string) (*model.AccessTokenClaims, error)
}

type Products interface{}

type Buckets interface{}

type Service struct {
	Authorization
	Products
	Buckets
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthSerice(repos.Authorization),
	}
}
