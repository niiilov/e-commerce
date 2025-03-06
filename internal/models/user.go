package models

type User struct {
	Id       int    `json: 'user_id'`
	Username string `json: 'username'`
	Email    string `json: 'email'`
	Phone    string `json: 'phone'`
	Password string `json: 'pass'`
}
