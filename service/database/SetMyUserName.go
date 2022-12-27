package database

import "errors"

func (db *appdbimpl) SetMyUserName(username string, new_username string) error {
	exists1, err := db.Exists(username)
	if err != nil {
		return err
	}
	if !exists1 {
		return errors.New("user does not exist")
	}

	query := "UPDATE users SET username = ? WHERE username = ?"
	_, err = db.c.Exec(query, new_username, username)
	if err != nil {
		return err
	}

	query = "UPDATE Follows SET followed = ? WHERE followed = ?"
	_, err = db.c.Exec(query, new_username, username)
	if err != nil {
		return err
	}
	query = "UPDATE Follows SET follower = ? WHERE follower = ?"
	_, err = db.c.Exec(query, new_username, username)
	if err != nil {
		return err
	}

	query = "UPDATE ban SET banner = ? WHERE banner = ?"
	_, err = db.c.Exec(query, new_username, username)
	if err != nil {
		return err
	}

	query = "UPDATE ban SET banned = ? WHERE banned = ?"
	_, err = db.c.Exec(query, new_username, username)
	if err != nil {
		return err
	}

	query = "UPDATE tokens SET username = ? WHERE username = ?"
	_, err = db.c.Exec(query, new_username, username)
	if err != nil {
		return err
	}

	query = "UPDATE photos SET username = ? WHERE username = ?"
	_, err = db.c.Exec(query, new_username, username)
	if err != nil {
		return err
	}

	query = "UPDATE comments SET username = ? WHERE username = ?"
	_, err = db.c.Exec(query, new_username, username)
	if err != nil {
		return err
	}

	query = "UPDATE likes SET username = ? WHERE username = ?"
	_, err = db.c.Exec(query, new_username, username)
	if err != nil {
		return err
	}
	return nil
}
