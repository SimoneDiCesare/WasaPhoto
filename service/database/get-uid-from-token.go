package database

import (
	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func (db *appdbimpl) GetUserIdFromToken(token string) (string, error) {
	var uid string
	uidError := db.c.QueryRow(queries.GetUseridFromToken, token).Scan(&uid)
	if uidError != nil {
		return "", uidError
	}
	return uid, nil
}
