package database

import (
	"database/sql"
	"errors"

	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
)

func (db *appdbimpl) SearchUidByToken(token string) (uid string, err error) {
	err = db.c.QueryRow(GetUidByToken, token).Scan(&uid)
	if errors.Is(err, sql.ErrNoRows) {
		db.logger.Debugf("No match found for token: %s", token)
		return "", err
	}
	return uid, nil
}

func (db *appdbimpl) SearchUsersByName(uid string, username string) (users []schema.SimpleUserData, err error) {
	rows, searchError := db.c.Query(SearchUsersByName, uid, username)
	if searchError != nil {
		db.logger.Debug("Search Error")
		return nil, searchError
	}
	defer func() {
		closeError := rows.Close()
		if closeError != nil && err != nil {
			db.logger.Debug("Close Error")
			err = closeError
		}
	}()
	for rows.Next() {
		var user schema.SimpleUserData
		if err := rows.Scan(&user.Uid, &user.Username); err != nil {
			db.logger.Debug("Scan Error")
			return nil, err
		}
		users = append(users, user)
	}
	rowsError := rows.Err()
	if rowsError != nil {
		db.logger.Debug("Rows Error")
		return nil, rowsError
	}
	return users, nil
}
