package database

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

func (db *appdbimpl) DoLogin(username string) (string, error) {
	token := fmt.Sprintf("%x", sha1.Sum([]byte(username)))
	_, err := db.c.Exec(`INSERT INTO users (username, token) VALUES (?, ?) ON CONFLICT DO NOTHING`, username, token)
	if err != nil {
		return "", err
	}
	err = db.c.QueryRow(`SELECT token FROM users WHERE username = ?`, username).Scan(&token)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (db *appdbimpl) FollowUser(token string, otheruserid string) error {
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

func (db *appdbimpl) UnfollowUser(token string, otheruserid string) error {
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

func (db *appdbimpl) BanUser(token string, otheruserid string) error {
	query := "INSERT INTO banned (token, banned) VALUES (?, ?)"
	_, err := db.c.Exec(query, token, otheruserid)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) UnbanUser(token string, otheruserid string) error {
	query := "DELETE FROM banned WHERE token = ? AND banned = ?"
	_, err := db.c.Exec(query, token, otheruserid)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) SetMyUserName(token string, new_username string) (string, error) {
	query := "UPDATE users SET username = ? , token = ? WHERE token = ?"
	new_token := fmt.Sprintf("%x", sha1.Sum([]byte(new_username)))
	_, err := db.c.Exec(query, new_username, new_token, token)
	if err != nil {
		return "", err
	}
	return new_token, nil
}

func (db *appdbimpl) GetUserProfile(token string, userid string) (User, error) {
	var user User
	var Follows string
	var following string
	query := "SELECT username, follows,following  FROM users WHERE username = ?"
	err := db.c.QueryRow(query, userid).Scan(&user.Username, &Follows, &following)
	if err != nil {
		return user, err
	}
	user.Follows = strings.Split(Follows, ",")
	user.Following = strings.Split(following, ",")
	query = "SELECT COUNT(*) FROM photos WHERE token = (SELECT token FROM users WHERE username = ?)"
	err = db.c.QueryRow(query, userid).Scan(&user.Nphotos)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *appdbimpl) GetUsers(token string, userid string) ([]string, error) {
	query := "SELECT username from users WHERE username LIKE ? "
	userid += "%"
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

func (db *appdbimpl) GetMyStream(token string) ([]Photo, error) {
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

func (db *appdbimpl) GetBanned(username string) ([]string, error) {
	query := "SELECT otheruserid FROM bans WHERE token = (SELECT token FROM users WHERE username = ?)"
	rows, err := db.c.Query(query, username)
	if err != nil {
		return nil, err
	}
	var banned []string
	for rows.Next() {
		var user string
		err = rows.Scan(&user)
		if err != nil {
			return nil, err
		}
		banned = append(banned, user)
	}
	return banned, nil
}
