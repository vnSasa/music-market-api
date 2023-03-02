package repository

import (
	"fmt"
	model "github.com/vnSasa/music-market-api/model"
	"database/sql"
	"errors"
)

type AuthDB struct {
	db *sql.DB
}

func NewAuthDB(db *sql.DB) *AuthDB {
	return &AuthDB{db: db}
}

func (r *AuthDB) CreateUser(user model.User) error {
	query := fmt.Sprintf("INSERT INTO %s (login, first_name, last_name, password) "+
			"VALUES (?, ?, ?, ?)", userTable)
	_, err := r.db.Exec(query, user.Login, user.FirstName, user.LastName, user.Password)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}