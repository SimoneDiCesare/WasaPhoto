package database

import (
	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
)

func (db *appdbimpl) GetPostComments(pid string) (comments []schema.PostComment, err error) {
	rows, searchError := db.c.Query(GetPostComments, pid)
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
		var comment schema.PostComment
		if err := rows.Scan(&comment.Cid, &comment.Pid, &comment.Author.Uid, &comment.Author.Username, &comment.Text); err != nil {
			db.logger.Debug("Scan Error")
			return nil, err
		}
		comments = append(comments, comment)
	}
	rowsError := rows.Err()
	if rowsError != nil {
		db.logger.Debug("Rows Error")
		return nil, rowsError
	}
	return comments, nil
}
