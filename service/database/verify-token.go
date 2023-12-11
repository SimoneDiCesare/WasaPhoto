package database

import (
	"database/sql"
	"errors"

	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func (db *appdbimpl) VerifyToken(token string) error {
	var token_found string
	err := db.c.QueryRow(queries.GetUserToken, token).Scan(&token_found)
	if err == sql.ErrNoRows {
		return errors.New("Invalid Authorization Token")
	}
	return nil
}
