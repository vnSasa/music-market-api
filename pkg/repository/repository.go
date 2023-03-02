package repository

import (
	"database/sql"
	model "github.com/vnSasa/music-market-api/model"
)

type Authorization interface {
	CreateUser(user model.User) error
}

type Products interface {

}

type Buckets interface {
	
}

type Repository struct {
	Authorization
	Products
	Buckets
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization:	NewAuthDB(db),
	}
}