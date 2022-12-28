package database

import (
	"errors"
	"hash/fnv"
	"regexp"
	"strconv"
	"time"
)

func (db *appdbimpl) UploadPhoto(username string, photo string) (string, error) {
	exists, err := db.Exists(username)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errors.New("user does not exist")
	}

	query := `INSERT INTO photos (id,username,photo, date) VALUES (?,?,?, ?)`
	//check if photo is a valid base64 string
	matched, err := regexp.MatchString(`data:image\/[^;]+;base64[^"]+`, photo)
	if err != nil {
		return "", err
	}
	if !matched {
		return "", errors.New("invalid photo")
	}
	//generate id
	h := fnv.New32a()
	h.Write([]byte(username + photo))
	id := strconv.Itoa(int(h.Sum32()))
	//insert into db
	_, err = db.c.Exec(query, id, username, photo, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return "", err
	}

	//update nphotos
	query = `UPDATE users SET nphotos = nphotos + 1 WHERE username = ?`
	_, err = db.c.Exec(query, username)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (db *appdbimpl) GetUserPhotos(username string) ([]string, error) {
	exists, err := db.Exists(username)
	if err != nil {
		return []string{}, err
	}
	if !exists {
		return []string{}, errors.New("user does not exist")
	}

	query := `SELECT id FROM photos WHERE username = ?`
	rows, err := db.c.Query(query, username)
	if err != nil {
		return nil, err
	}
	var photos []string
	for rows.Next() {
		var photo string
		err = rows.Scan(&photo)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}
	return photos, nil
}

func (db *appdbimpl) DeletePhoto(photoid string) error {
	//get photo owner username
	var username string
	query := `SELECT username FROM photos WHERE id = ?`
	err := db.c.QueryRow(query, photoid).Scan(&username)
	if err != nil {
		return err
	}
	query = `DELETE FROM photos WHERE id = ?`
	_, err = db.c.Exec(query, photoid)
	if err != nil {
		return err
	}
	//update nphotos
	query = `UPDATE users SET nphotos = nphotos - 1 WHERE username = ?`
	_, err = db.c.Exec(query, username)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetPhoto(photoid string) (Photo, error) {
	query := `SELECT photo FROM photos WHERE id = ?`
	var photo Photo
	err := db.c.QueryRow(query, photoid).Scan(&photo.Photo)
	photo.Id = photoid
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
