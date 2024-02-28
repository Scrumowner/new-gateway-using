package service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"user_service/infrastructure/tools/crypt"
	"user_service/internal/models"
	"user_service/internal/modules/user/storage"
)

type UserService struct {
	storage *storage.UserStorage
	token   *crypt.TokenManager
}

func NewUserSerivce(db *sqlx.DB) *UserService {
	return &UserService{
		storage: storage.NewUserStorage(db),
	}
}

func (u *UserService) GetUser(user *models.User) (string, error) {
	userinDb, err := u.storage.GetUser(user)
	if err != nil {
		return "", err
	}
	token, err := u.token.CreateToken(userinDb.Username, userinDb.Email, userinDb.Password)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Bearer %s", token), nil
}

func (u *UserService) SetUser(user *models.User) error {
	err := u.storage.SetUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) Auth(token string) bool {
	userFromPayload, err := u.token.Parse(token)
	if err != nil {
		return false
	}
	_, err = u.storage.GetUser(&models.User{
		Username: userFromPayload.Username,
		Email:    userFromPayload.Email,
		Password: userFromPayload.Password,
	})
	if err != nil {
		return false
	}
	return true

}
