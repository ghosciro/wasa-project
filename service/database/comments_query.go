package database

func (db *appdbimpl) CommentPhoto(token string, photoid string, comment string) (int, error) {
	query := "INSERT INTO comments (token, photoid, comment) VALUES (?, ?, ?)"
	result, err := db.c.Exec(query, token, photoid, comment)
	if err != nil {
		return -1, err
	}
	resultid, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(resultid), nil
}

func (db *appdbimpl) UncommentPhoto(token string, photoid string, comment int) error {
	query := "DELETE FROM comments WHERE token = ? AND photoid = ? AND comment = ?"
	_, err := db.c.Exec(query, token, photoid, comment)
	if err != nil {
		return err
	}
	return nil
}
