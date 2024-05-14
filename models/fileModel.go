package models

type File struct {
	FileID   int64  `json:"file_id"`
	Filename string `json:"filename"`
}

func InsertFile(fileName string, userid int64) (int64, error) {
	res, err := db.Exec("INSERT INTO file VALUES(DEFAULT, ?, ?)", fileName, userid)

	fileID, _ := res.LastInsertId()
	return fileID, err
}

func GetFile(fileID int64, userID int64) (File, error) {
	fileData := File{}

	row := db.QueryRow("SELECT * FROM file WHERE file_id = ? AND user_id = ?", fileID, userID)

	err := row.Scan(&fileData.FileID, &fileData.Filename, &userID)

	return fileData, err
}

func GetAllFiles(userID int64) ([]File, error) {
	filesData := []File{}

	rows, err := db.Query("SELECT * FROM file WHERE user_id = ?", userID)
	for rows.Next() {
		fileData := File{}
		rows.Scan(&fileData.FileID, &fileData.Filename, &userID)
		filesData = append(filesData, fileData)
	}

	return filesData, err
}
