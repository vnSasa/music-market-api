package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	model "github.com/vnSasa/music-market-api/model"
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
		return err
	}

	return nil
}

func (r *AuthDB) GetUser(login, password string) (int, error) {
	// CHECK INPUT DATA
	var pwd string
	confirmPassword := fmt.Sprintf("SELECT password FROM %s WHERE login = ?", userTable)
	row := r.db.QueryRow(confirmPassword, login)
	err := row.Scan(&pwd)
	if err != nil {
		return 0, errors.New("password not found")
	}
	if strings.Compare(pwd, password) != 0 {
		return 0, errors.New("error input password")
	}

	// SEARCH ID
	var id int
	searchID := fmt.Sprintf("SELECT id FROM %s WHERE login = ?", userTable)
	rowID := r.db.QueryRow(searchID, login)
	err = rowID.Scan(&id)
	if err != nil {
		return 0, errors.New("something went wrong when write id")
	}

	return id, nil
}
