package models

type File struct {
	FileID   string
	Filename string
	FileHash string
}

func InsertFile(file File, userid string) (string, error) {
	_, err := db.Exec("INSERT INTO file VALUES(DEFAULT, ?, ?, ?)", file.Filename, file.FileHash, userid)
	row := db.QueryRow("SELECT file_id FROM file WHERE file_hash = ?", file.FileHash)

	row.Scan(&file.FileID)

	return file.FileID, err
}
