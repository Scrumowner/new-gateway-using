package storage

import (
	"github.com/jmoiron/sqlx"
	"user_service/internal/models"
)

type UserStorage struct {
	db *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (u *UserStorage) GetUser(user *models.User) (*models.User, error) {
	res := &models.User{}
	query := "SELECT * FROM users WHERE (username=$1, email=$2, password=$3)"
	err := u.db.Get(res, query, user.Username, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (u *UserStorage) SetUser(user *models.User) error {
	query := "INSERT INTO users (username, email, password) VALUES ($1,$2,$3)"
	_, err := u.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
