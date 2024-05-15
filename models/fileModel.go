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

	row := db.QueryRow(`
		SELECT file_id, filename, user_id FROM file
			WHERE file_id=? AND user_id=?
				UNION
		SELECT shared_file.file_id, file.filename, shared_file.user_id FROM shared_file
				INNER JOIN file
			ON shared_file.file_id = file.file_id
			HAVING shared_file.file_id=? AND shared_file.user_id=?
	`, fileID, userID, fileID, userID)

	err := row.Scan(&fileData.FileID, &fileData.Filename, &userID)

	return fileData, err
}

func GetAllFiles(userID int64) ([]File, error) {
	filesData := []File{}

	rows, err := db.Query(`
		SELECT file_id, filename, user_id FROM file
			WHERE user_id=?
				UNION
		SELECT shared_file.file_id, file.filename, shared_file.user_id FROM shared_file
				INNER JOIN file
			ON shared_file.file_id = file.file_id
			HAVING shared_file.user_id=?
	`, userID, userID)

	for rows.Next() {
		fileData := File{}
		rows.Scan(&fileData.FileID, &fileData.Filename, &userID)
		filesData = append(filesData, fileData)
	}

	return filesData, err
}

func InsertSharedFile(fileID int64, targetUserID int64) error {
	_, err := db.Exec("INSERT INTO shared_file VALUES(?, ?)", fileID, targetUserID)

	return err
}
