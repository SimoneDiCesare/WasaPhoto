package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) VerifyToken(token string) (err error) {
	var token_found string
	err = db.c.QueryRow("SELECT token FROM users WHERE token=$1", token).Scan(&token_found)
	if err == sql.ErrNoRows {
		// User does not exist, create a new one
		return errors.New("Invalid Authorization Token")
	}
	return nil
}
