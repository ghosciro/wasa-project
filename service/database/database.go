/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	Username  string
	Followers []string
	Follows   []string
	Nphotos   int
}
type Photo struct {
	Id    int64
	Photo string
	Date  string
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	BanUser(token string, otheruserid string) error
	CommentPhoto(token string, photoid string, comment string) (int, error)
	DeletePhoto(token string, photoid string) error
	DoLogin(username string) (string, error)
	FollowUser(token string, otheruserid string) error
	GetMyStream(token string) ([]Photo, error)
	GetPhoto(token string, id string) (Photo, error)
	GetUserProfile(token string, userid string) (User, error)
	GetUsers(token string, username string) ([]string, error)
	LikePhoto(token string, photoid string) error
	SetMyUserName(token string, new_username string) error
	UnbanUser(token string, otheruserid string) error
	UncommentPhoto(token string, photoid string, comment int) error
	UnfollowUser(token string, otheruserid string) error
	UnlikePhoto(token string, photoid string) error
	UploadPhoto(token string, photo string) (string, error)
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	var err error
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}
	//create user table
	usertable := "CREATE TABLE IF NOT EXISTS Users (token TEXT NOT NULL PRIMARY KEY, username TEXT NOT NULL PRIMARY KEY, follows TEXT, following TEXT);"
	_, err = db.Exec(usertable)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	phototable := "CREATE TABLE IF NOT EXISTS photos (id TEXT NOT NULL PRIMARY KEY, FOREIGN KEY (token) REFERENCES Users(token), photo TEXT, date TEXT);"
	_, err = db.Exec(phototable)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	//create comments table
	commenttable := "CREATE TABLE IF NOT EXISTS comments (id INTEGER NOT NULL PRIMARY KEY, FOREIGN KEY (token) REFERENCES Users(token), FOREIGN KEY (photoid) REFERENCES photos(id), comment NOT NULL TEXT PRIMARY KEY);"
	_, err = db.Exec(commenttable)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	//create ban table
	bantable := "CREATE TABLE IF NOT EXISTS ban (FOREIGN KEY (token) REFERENCES Users(token), FOREIGN KEY (otheruserid) REFERENCES Users(token));"
	_, err = db.Exec(bantable)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	//create like table
	liketable := "CREATE TABLE IF NOT EXISTS likes (id INTEGER NOT NULL PRIMARY KEY, FOREIGN KEY (token) REFERENCES Users(token), FOREIGN KEY (photoid) REFERENCES photos(id));"
	_, err = db.Exec(liketable)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
