package database

import (
	"errors"

	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
)

func (db *appdbimpl) ChangeUserName(uid string, username string) error {
	var tmp1, tmp2, tmp3 string
	queryError := db.c.QueryRow(GetUserByName, username).Scan(&tmp1, &tmp2, &tmp3)
	db.logger.Debugf("Searching username: %s -> %e", username, queryError)
	if errors.Is(queryError, nil) {
		// Username already existing
		db.logger.Debug("Can't change username to one already in use.")
		return schema.ErrExistingUsername
	}
	result, err := db.c.Exec(UpdateUserName, username, uid)
	if err != nil {
		db.logger.Debugf("Error on db Query: %e", err)
		return err
	}
	count, err := result.RowsAffected()
	if count == 0 || err != nil {
		db.logger.Debugf("No rows affected: %d, %e", count, err)
		return err
	}
	return nil
}
