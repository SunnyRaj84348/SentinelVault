package models

type User struct {
	UserID   string
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func InsertUser(user User) error {
	_, err := db.Exec(`INSERT INTO user VALUES(DEFAULT, ?, ?)`, user.Username, user.Password)

	return err
}
