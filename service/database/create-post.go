package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) CreatePost(uid string) (pid string, err error) {
	triesCount := 0
	var foundId string
	for {
		pid, err = generateRandomHex()
		if err != nil {
			db.logger.Errorf("Error generating hex: %e", err)
			return "", err
		}
		queryError := db.c.QueryRow(GetPostIdFromId, pid).Scan(&foundId)
		if errors.Is(queryError, sql.ErrNoRows) {
			break
		}
		triesCount = triesCount + 1
		if triesCount > 10 {
			return "", errors.New("exceeded tries for creating pid")
		}
	}
	_, queryError := db.c.Exec(CreatePost, pid, uid)
	if queryError != nil {
		db.logger.Errorf("Error on Creating Post on db: %e", queryError)
		return "", queryError
	}
	return pid, nil
}
