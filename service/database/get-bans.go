package database

import (
	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
)

func (db *appdbimpl) GetBans(uid string) (bans []schema.SimpleUserData, err error) {
	rows, searchError := db.c.Query(GetBans, uid)
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
		var ban schema.SimpleUserData
		if err := rows.Scan(&ban.Uid, &ban.Username); err != nil {
			db.logger.Debug("Scan Error")
			return nil, err
		}
		bans = append(bans, ban)
	}
	rowsError := rows.Err()
	if rowsError != nil {
		db.logger.Debug("Rows Error")
		return nil, rowsError
	}
	return bans, nil
}
