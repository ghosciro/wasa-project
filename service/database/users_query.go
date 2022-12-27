package database

import (
	"errors"
	"strings"
)

func (db *appdbimpl) DoLogin(username string) (string, error) {
	query := "INSERT INTO users (username,nphotos) VALUES (?, ?) on conflict do nothing"
	_, err := db.c.Exec(query, username, 0)
	if err != nil {
		return "", err
	}
	row, err := db.c.Query("select token from tokens where username = ?", username)
	//check if token exists
	//if exists, return error "already logged in"
	//else generate token and insert into db
	if err != nil {
		print("error here 2")
		return "", err
	}

	if row == nil {
		return "", errors.New("already logged in")
	}

	//generate 20 random character
	token := Generate_random_string(30)
	query = "INSERT INTO tokens(username, token) VALUES (?, ?)"
	_, err = db.c.Exec(query, username, token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (db *appdbimpl) FollowUser(username string, otherusername string) error {
	exists1, err := db.Exists(otherusername)
	exists2, err2 := db.Exists(username)
	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}
	if !exists1 || !exists2 {
		return errors.New("user does not exist")
	}

	query := "INSERT INTO Follows(followed, follower) VALUES (?, ?) on conflict do nothing"
	_, err = db.c.Exec(query, otherusername, username)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) UnfollowUser(username string, otherusername string) error {
	exists1, err := db.Exists(otherusername)
	exists2, err2 := db.Exists(username)
	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}
	if !exists1 || !exists2 {
		return errors.New("user does not exist")
	}
	query := "DELETE FROM Follows WHERE followed = ? AND follower = ?"
	_, err = db.c.Exec(query, otherusername, username)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) BanUser(username string, otherusername string) error {
	exists1, err := db.Exists(otherusername)
	exists2, err2 := db.Exists(username)
	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}
	if !exists1 || !exists2 {
		return errors.New("user does not exist")
	}

	query := "INSERT INTO ban (banner, banned) VALUES (?, ?)"
	_, err = db.c.Exec(query, username, otherusername)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) UnbanUser(username string, otherusername string) error {

	exists1, err := db.Exists(otherusername)
	exists2, err2 := db.Exists(username)
	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}
	if !exists1 || !exists2 {
		return errors.New("user does not exist")
	}

	query := "DELETE FROM ban WHERE banner = ? AND banned = ?"
	_, err = db.c.Exec(query, username, otherusername)
	if err != nil {
		return err
	}
	return nil
}


func (db *appdbimpl) GetUserProfile(username string) (User, error) {
	exists, err := db.Exists(username)
	if err != nil {
		return User{}, err
	}
	if !exists {
		return User{}, errors.New("user does not exist")
	}

	var user User

	query := "SELECT username, nphotos  FROM users WHERE username = ?"
	err = db.c.QueryRow(query, username).Scan(&user.Username, &user.Nphotos)
	if err != nil {
		return user, err
	}

	query = "SELECT follower FROM Follows WHERE followed = ?"
	row, err := db.c.Query(query, username)
	if err != nil {
		return user, err
	}
	for row.Next() {
		var follower string
		err = row.Scan(&follower)
		if err != nil {
			return user, err
		}
		user.Follower = append(user.Follower, follower)
	}
	query = "SELECT followed FROM Follows WHERE follower = ?"
	row, err = db.c.Query(query, username)
	if err != nil {
		return user, err
	}
	for row.Next() {
		var followed string
		err = row.Scan(&followed)
		if err != nil {
			return user, err
		}
		user.Follows = append(user.Follows, followed)
	}
	return user, nil
}

func (db *appdbimpl) GetUsers(userid string) ([]string, error) {
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

func (db *appdbimpl) GetMyStream(username string) ([]Photo, error) {
	var follow string
	var photos []Photo
	query := "SELECT follows FROM users WHERE username = ?"
	err := db.c.QueryRow(query, username).Scan(&follow)
	if err != nil {
		return photos, err
	}
	followers := strings.Split(follow, ",")
	for _, follower := range followers {
		query := "SELECT id, photo, date FROM photos WHERE username = (SELECT username FROM users WHERE username = ? )"
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
	query := "SELECT banned FROM ban WHERE banner = ?"
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
