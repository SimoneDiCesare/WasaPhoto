package database

import (
	"errors"

	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func (db *appdbimpl) GetPostLikes(pid string) (likes []Like, err error) {
	rows, retrievingError := db.c.Query(queries.GetPostLikes, pid)
	if retrievingError != nil {
		return nil, errors.New("Error retrieving comments")
	}
	defer func() {
		closeErr := rows.Close()
		if err == nil {
			err = closeErr
		}
	}()
	for rows.Next() {
		var like Like
		if err := rows.Scan(&like.Pid, &like.Uid); err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	rowsErr := rows.Err()
	if rowsErr != nil {
		return nil, rowsErr
	}
	return likes, nil
}
