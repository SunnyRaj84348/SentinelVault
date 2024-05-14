package models

type User struct {
	UserID   int64
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func InsertUser(user User) error {
	_, err := db.Exec(`INSERT INTO user VALUES(DEFAULT, ?, ?)`, user.Username, user.Password)

	return err
}

func GetUser(username string) (User, error) {
	user := User{}

	row := db.QueryRow(`SELECT * FROM user WHERE username = ?`, username)
	if row.Err() != nil {
		return user, row.Err()
	}

	row.Scan(&user.UserID, &user.Username, &user.Password)

	return user, nil
}
