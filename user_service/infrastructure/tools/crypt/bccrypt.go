package crypt

import "golang.org/x/crypto/bcrypt"

// return hahsed password string and error
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// compare password and hash and return bool
func CheckPasswrod(password string, hash string) bool {
	isValid := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if isValid != nil {
		return false
	}
	return true
}
