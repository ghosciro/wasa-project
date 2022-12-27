package database

import (
	"crypto/sha1"
	"errors"
	"regexp"
	"time"
)

func (db *appdbimpl) UploadPhoto(username string, photo string) (string, error) {
	query := `INSERT INTO photos (id,photo, date) VALUES (?,?, ?)`
	//check if photo is a valid base64 string
	matched, err := regexp.MatchString(`data:image\/[^;]+;base64[^"]+`, photo)
	if err != nil {
		return "", err
	}
	if !matched {
		return "", errors.New("invalid base64 string")
	}
	//generate id
	hash := sha1.Sum([]byte(username + time.Now().String()))
	id := string(hash[:])

	//insert into db
	_, err = db.c.Exec(query, id, photo, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return "", err
	}
	return id, nil
}

func (db *appdbimpl) DeletePhoto(photoid string) error {
	query := `DELETE FROM photos WHERE id = ?`
	_, err := db.c.Exec(query, photoid)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetPhoto(photoid string) (Photo, error) {
	query := `SELECT photo FROM photos WHERE id = ?`
	var photo Photo
	err := db.c.QueryRow(query, photoid).Scan(&photo.Photo)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (db *appdbimpl) LikePhoto(username string, photoid string) error {
	query := `INSERT INTO likes (username, photoid) VALUES (?,?)`
	_, err := db.c.Exec(query, username, photoid)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) UnlikePhoto(username string, photoid string) error {
	query := `DELETE FROM likes WHERE(username = ? AND photoid = ?`
	_, err := db.c.Exec(query, username, photoid)
	if err != nil {
		return err
	}
	return nil
}
