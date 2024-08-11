package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) LikePost(uid string, pid string) error {
	var tmpId string
	queryError := db.c.QueryRow(CheckLike, uid, pid).Scan(tmpId)
	if errors.Is(queryError, sql.ErrNoRows) {
		// No like present
		_, queryError := db.c.Exec(LikePost, uid, pid)
		if queryError != nil {
			db.logger.Debugf("Error liking: %e", queryError)
			return queryError
		}
		return nil
	}
	return nil
}

func (db *appdbimpl) UnlikePost(uid string, pid string) error {
	_, queryError := db.c.Exec(UnlikePost, uid, pid)
	if queryError != nil {
		db.logger.Debugf("Error unliking: %e", queryError)
		return queryError
	}
	return nil
}
