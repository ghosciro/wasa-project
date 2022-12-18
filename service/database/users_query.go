package database

import (
	"strings"
)

func (db *appdbimpl) doLogin(username string) (string, error) {
	query := `SELECT token FROM users WHERE username = ?`
	var token string
	err := db.c.QueryRow(query, username).Scan(&token)
	if err != nil {
		return token, err
	}
	return token, nil
}

func (db *appdbimpl) followUser(token string, otheruserid string) error {
	query := "SELECT  follows FROM users WHERE  token = ?"
	var follows string
	err := db.c.QueryRow(query, token).Scan(&follows)
	if err != nil {
		return err
	}
	follows += otheruserid + ","
	query = "update users set follows = ? where token = ?"
	_, err = db.c.Exec(query, follows, token)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) unfollowUser(token string, otheruserid string) error {
	query := "SELECT follows fFROMusers  WHERE token = ?"
	var follows string
	err := db.c.QueryRow(query, token).Scan(&follows)
	if err != nil {
		return err
	}
	follows = strings.Replace(follows, otheruserid+",", "", 1)
	query = "update users set follows = ? where token = ?"
	_, err = db.c.Exec(query, follows, token)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) banUser(token string, otheruserid string) error {
	query := "INSERT INTO banned (token, banned) VALUES (?, ?)"
	_, err := db.c.Exec(query, token, otheruserid)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) unbanUser(token string, otheruserid string) error {
	query := "DELETE FROM banned WHERE token = ? AND banned = ?"
	_, err := db.c.Exec(query, token, otheruserid)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) setMyUserName(token string, new_username string) error {
	query := "UPDATE users SET username = ? WHERE token = ?"
	_, err := db.c.Exec(query, new_username, token)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) getUserProfile(token string, userid string) (User, error) {
	var user User
	query := "SELECT username, follows FROM users WHERE username = ?"
	err := db.c.QueryRow(query, userid).Scan(&user.Username, &user.Follows)
	if err != nil {
		return user, err
	}
	query = "SELECT COUNT(*) FROM photos WHERE token = (SELECT token FROM users WHERE username = ?)"
	err = db.c.QueryRow(query, userid).Scan(&user.Nphotos)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *appdbimpl) getUsers(token string, userid string) ([]string, error) {
	query := "SELECT username from users WHERE username LIKE ?%"
	rows, err := db.c.Query(query, userid)
	if err != nil {
		return nil, err
	}
	var users []string
	for rows.Next() {
		var user string
		err = rows.Scan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (db *appdbimpl) getMyStream(token string) ([]Photo, error) {
	var follow string
	var photos []Photo
	query := "SELECT follows FROM users WHERE token = ?"
	err := db.c.QueryRow(query, token).Scan(&follow)
	if err != nil {
		return photos, err
	}
	followers := strings.Split(follow, ",")
	for _, follower := range followers {
		query := "SELECT id, photo, date FROM photos WHERE token = (SELECT token FROM users WHERE username = ? )"
		rows, err := db.c.Query(query, follower)
		if err != nil {
			return photos, err
		}
		for rows.Next() {
			var photo Photo
			err = rows.Scan(&photo.Id, &photo.Photo, &photo.Date)
			if err != nil {
				return photos, err
			}
			photos = append(photos, photo)
		}
	}
	//order photos by date
	for i := 0; i < len(photos); i++ {
		for j := i + 1; j < len(photos); j++ {
			if photos[i].Date < photos[j].Date {
				photos[i], photos[j] = photos[j], photos[i]
			}
		}
	}
	return photos, nil

}
