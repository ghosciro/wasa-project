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

	//  Start Database
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
	Username string
	Follows  []string
	Follower []string
	Nphotos  int
}
type Photo struct {
	Id    string
	Photo string
	Date  string
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	BanUser(username string, otherusername string) error
	CommentPhoto(username string, photoid string, comment string) (int, error)
	DeletePhoto(photoid string) error
	DoLogin(username string) (string, error)
	FollowUser(username string, otherusername string) error
	GetMyStream(username string) ([]Photo, error)
	GetPhoto(id string) (Photo, error)
	GetUserProfile(userid string) (User, error)
	GetUsers(username string) ([]string, error)
	LikePhoto(username string, photoid string) error
	SetMyUserName(username string, new_username string) error
	UnbanUser(username string, otherusername string) error
	UncommentPhoto(username string, photoid string, comment int) error
	UnfollowUser(username string, otherusername string) error
	UnlikePhoto(username string, photoid string) error
	UploadPhoto(username string, photo string) (string, error)
	GetBanned(username string) ([]string, error)
	Exists(username string) (bool, error)
	GetUserPhotos(username string) ([]string, error)
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
	// create user table
	usertable := `
		CREATE TABLE IF NOT EXISTS Users
		(
		username TEXT NOT NULL ,
		nphotos INTEGER DEFAULT 0,
		CONSTRAINT userPK PRIMARY KEY (username));`

	_, err = db.Exec(usertable)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	// create follow table
	followtable := `
		CREATE TABLE IF NOT EXISTS Follows
		(
		followed TEXT NOT NULL default "",
		follower TEXT NOT NULL default "", 
		CONSTRAINT followPK PRIMARY KEY (followed, follower),
		FOREIGN KEY (followed) REFERENCES Users(username),
		FOREIGN KEY (follower) REFERENCES Users(username));`
	_, err = db.Exec(followtable)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	// create photo table

	phototable := `
		CREATE TABLE IF NOT EXISTS photos  
		(
		 id TEXT NOT NULL , 
		 username TEXT NOT NULL,
		 photo TEXT DEFAULT "",
		 date TEXT DEFAULT "",			
		 FOREIGN KEY (username) REFERENCES Users(username)
		 CONSTRAINT photoPK PRIMARY KEY (id, username)
		);
		`
	_, err = db.Exec(phototable)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	// create comments table
	commenttable := `
	CREATE TABLE IF NOT EXISTS comments
	(id INTEGER NOT NULL ,
	comment TEXT NOT NULL,
	photoid TEXT NOT NULL,
	username TEXT NOT NULL,
	FOREIGN KEY (username) REFERENCES Users(username) ,
	FOREIGN KEY (photoid) REFERENCES photos(id),
	CONSTRAINT commPK PRIMARY KEY (id))
	;`
	_, err = db.Exec(commenttable)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	// create ban table
	bantable := `
	CREATE TABLE IF NOT EXISTS ban
	 (
	  banner TEXT NOT NULL,
	  banned TEXT NOT NULL,
	  FOREIGN KEY (banner) REFERENCES Users(username),
	  FOREIGN KEY (banned) REFERENCES Users(username),
	  CONSTRAINT banPK PRIMARY KEY (banner, banned)
	  );`
	_, err = db.Exec(bantable)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	// create like table
	liketable := `CREATE TABLE IF NOT EXISTS likes
		(id INTEGER NOT NULL,
		username TEXT NOT NULL ,
		photoid TEXT NOT NULL,
		FOREIGN KEY (photoid) REFERENCES photos(id),
		FOREIGN KEY (username) REFERENCES Users(username),
		CONSTRAINT likePK PRIMARY KEY (username, photoid)
		);`
	_, err = db.Exec(liketable)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	tokentable := `Create table if not exists tokens
		(username TEXT NOT NULL,
		token TEXT NOT NULL,
		CONSTRAINT tokenPK PRIMARY KEY (username));`
	_, err = db.Exec(tokentable)
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
