package models

type File struct {
	FileID   string `json:"file_id"`
	Filename string `json:"filename"`
	FileHash string `json:"-"`
}

func InsertFile(file File, userid string) (string, error) {
	_, err := db.Exec("INSERT INTO file VALUES(DEFAULT, ?, ?, ?)", file.Filename, file.FileHash, userid)
	row := db.QueryRow("SELECT file_id FROM file WHERE file_hash = ?", file.FileHash)

	row.Scan(&file.FileID)

	return file.FileID, err
}

func GetFile(fileID string, userID string) (File, error) {
	fileData := File{}

	row := db.QueryRow("SELECT * FROM file WHERE file_id = ? AND user_id = ?", fileID, userID)

	err := row.Scan(&fileData.FileID, &fileData.Filename, &fileData.FileHash, &userID)

	return fileData, err
}

func GetAllFiles(userID string) ([]File, error) {
	filesData := []File{}

	rows, err := db.Query("SELECT * FROM file WHERE user_id = ?", userID)
	for rows.Next() {
		fileData := File{}
		rows.Scan(&fileData.FileID, &fileData.Filename, &fileData.FileHash)
		filesData = append(filesData, fileData)
	}

	return filesData, err
}
