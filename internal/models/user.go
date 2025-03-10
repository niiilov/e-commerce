package models

type User struct {
	Id       int    `json: 'user_id'`
	Username string `json: 'username'`
	Email    string `json: 'email'`
	Password string `json: 'pass'`
}
