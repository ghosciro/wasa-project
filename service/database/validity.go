package database

import (
	"math/rand"
)

func (db *appdbimpl) Exists(username string) (bool, error) {
	query := "SELECT username FROM users WHERE username = ?"
	var user string
	err := db.c.QueryRow(query, username).Scan(&user)
	if err != nil {
		return false, err
	}
	if user == username {
		return true, nil
	}
	return false, nil
}

func Generate_random_string(length int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
