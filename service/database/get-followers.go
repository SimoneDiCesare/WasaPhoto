package database

import (
	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
)

func (db *appdbimpl) GetFollowers(uid string) (follows []schema.SimpleUserData, err error) {
	rows, searchError := db.c.Query(GetFollowers, uid)
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
		var follow schema.SimpleUserData
		if err := rows.Scan(&follow.Uid, &follow.Username); err != nil {
			db.logger.Debug("Scan Error")
			return nil, err
		}
		follows = append(follows, follow)
	}
	rowsError := rows.Err()
	if rowsError != nil {
		db.logger.Debug("Rows Error")
		return nil, rowsError
	}
	return follows, nil
}
