package database

import (
	"hash/fnv"
	"strconv"
)

func (db *appdbimpl) CommentPhoto(username string, photoid string, comment string) (int, error) {
	query := "INSERT INTO comments (id,username, photoid, comment) VALUES (?,?, ?, ?)"
	h := fnv.New32a()
	h.Write([]byte(username + photoid+comment))
	id := strconv.Itoa(int(h.Sum32()))

	result, err := db.c.Exec(query, id, username, photoid, comment)
	if err != nil {
		return -1, err
	}
	resultid, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(resultid), nil
}

func (db *appdbimpl) UncommentPhoto(username string, photoid string, comment_id int) error {
	query := "DELETE FROM comments WHERE username = ? AND photoid = ? AND id = ?"
	_, err := db.c.Exec(query, username, photoid, comment_id)
	if err != nil {
		return err
	}
	return nil
}
