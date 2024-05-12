package models

type User struct {
	UserID   string
	Username string `json:"username"`
	Password string `json:"password"`
}
