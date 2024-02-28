package models

import "fmt"

type User struct {
	Username string `json:"username" bd:"username"  bd_type:"text"`
	Email    string `json:"email" bd:"email" bd_type:"text"`
	Password string `json:"password" bd:"password" bd_type:"text"`
}

func (u *User) TableName() string {
	return "user"
}
func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetUsername(username string) error {
	if username == "" {
		return fmt.Errorf("Username should not be empty")
	}
	u.Username = username
	return nil
}

func (u *User) SetEmail(email string) error {
	if email == "" {
		return fmt.Errorf("Email should not be empty")
	}
	u.Email = email
	return nil
}

func (u *User) SetPassword(password string) error {
	if password == "" {
		return fmt.Errorf("Password should not be empty")
	}
	u.Password = password
	return nil
}
