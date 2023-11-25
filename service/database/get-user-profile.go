package database

import (
	"database/sql"
	"errors"

	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func (db *appdbimpl) GetUserProfile(uid string) (*User, error) {
	var user User
	err := db.c.QueryRow(queries.GetUserFromUid, uid).Scan(&user.Uid, &user.Username, &user.Biography, &user.Token)
	if err == sql.ErrNoRows {
		return nil, errors.New("No user found")
	}
	return &user, nil
}
