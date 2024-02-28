package crypt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserClaims struct {
	Username string `json:"username" bd:"username"  bd_type:"text"`
	Email    string `json:"email" bd:"email" bd_type:"text"`
	Password string `json:"password" bd:"password" bd_type:"text"`
	jwt.RegisteredClaims
}

type TokenManager struct {
	Key []byte
}

func NewTokenMnager(key []byte) *TokenManager {
	return &TokenManager{
		Key: key,
	}
}

func (t *TokenManager) CreateToken(username, email, password string) (string, error) {
	userClaims := &UserClaims{
		Username: username,
		Email:    email,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add((time.Hour * 24) * 30)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tkn, err := token.SignedString(t.Key)
	return tkn, err
}

func (t *TokenManager) Parse(token string) (*UserClaims, error) {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return t.Key, nil
	})
	if err != nil {
		return &UserClaims{}, err
	}
	if !tkn.Valid {
		return &UserClaims{}, fmt.Errorf("Token is invalid")
	}
	claims, ok := tkn.Claims.(jwt.MapClaims)
	if !ok {
		return &UserClaims{}, fmt.Errorf("error get user claims from token")
	}
	return &UserClaims{
		Username:         claims["username"].(string),
		Email:            claims["email"].(string),
		Password:         claims["email"].(string),
		RegisteredClaims: jwt.RegisteredClaims{},
	}, nil

}
