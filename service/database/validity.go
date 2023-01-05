package database

import (
	"math/rand"
)

func (db *appdbimpl) Exists(username string) (bool, error) {
	query := "SELECT username FROM users WHERE username = ?"
	var user string
	row := db.c.QueryRow(query, username)
	if row.Err() != nil {
		return false, row.Err()
	}
	err := row.Scan(&user)
	if err != nil {
		return false, err
	}
	if user == username {
		return true, nil
	}
	return false, nil
}

func (db *appdbimpl) Isnotbanned(username string, otherusername string) bool {
	query := "Select count( banned) from ban where banner = ? and banned = ?"
	var banned int
	row := db.c.QueryRow(query, otherusername, username)
	if row.Err() != nil {
		return false
	}
	err := row.Scan(&banned)
	if err != nil {
		return false
	}
	return banned == 0

}

func Generate_random_string(length int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
