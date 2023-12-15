package database

import (
	"errors"

	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func (db *appdbimpl) GetPostComment(cid string, pid string) (*Comment, error) {
	comment := &Comment{}
	retrievingError := db.c.QueryRow(queries.GetPostComment, cid, pid).Scan(&comment.Cid, &comment.Pid, &comment.Uid, &comment.Content)
	if retrievingError != nil {
		return nil, errors.New("Error retrieving comments")
	}
	return comment, nil
}
