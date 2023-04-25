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
	query := fmt.Sprintf("INSERT INTO %s (login, first_name, last_name, password, status) "+
		"VALUES (?, ?, ?, ?, ?)", userTable)
	_, err := r.db.Exec(query, user.Login, user.FirstName, user.LastName, user.Password, user.Status)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthDB) GetUserID(login, password string) (int, error) {
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

func (r *AuthDB) GetUserByID(id int) (*model.User, error) {
	var user model.User
	confirmUser := fmt.Sprintf("SELECT login, first_name, last_name, password, status FROM %s WHERE id = ?", userTable)
	row := r.db.QueryRow(confirmUser, id)
	err := row.Scan(&user.Login, &user.FirstName, &user.LastName, &user.Password, &user.Status)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *AuthDB) UpdateUserWithPassword(userID int, input model.User) error {
	query := fmt.Sprintf("UPDATE %s SET login=?, first_name=?, last_name=?, password=? WHERE id=?", userTable)

	_, err := r.db.Exec(query, input.Login, input.FirstName, input.LastName, input.Password, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthDB) UpdateUserWithoutPassword(userID int, input model.User) error {
	query := fmt.Sprintf("UPDATE %s SET login=?, first_name=?, last_name=? WHERE id=?", userTable)

	_, err := r.db.Exec(query, input.Login, input.FirstName, input.LastName, userID)
	if err != nil {
		return err
	}

	return nil
}
