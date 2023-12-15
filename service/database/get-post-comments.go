package database

import (
	"errors"

	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func (db *appdbimpl) GetPostComments(pid string) (comments []Comment, err error) {
	rows, retrievingError := db.c.Query(queries.GetPostComments, pid)
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
		var comment Comment
		if err := rows.Scan(&comment.Cid, &comment.Pid, &comment.Uid, &comment.Content); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	rowsErr := rows.Err()
	if rowsErr != nil {
		return nil, rowsErr
	}
	return comments, nil
}
