package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `json: 'user_id'`
	Username string `json: 'username'`
	Email    string `json: 'email'`
	Password string `json: 'pass'`
}

func (u *User) HashPassword() error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashPassword)

	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
