package database

import (
	"errors"

	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func (db *appdbimpl) DeletePostComment(token string, cid string, pid string) error {
	uid, uidError := db.GetUserIdFromToken(token)
	if uidError != nil {
		return uidError
	}
	comment, getError := db.GetPostComment(cid, pid)
	if getError != nil {
		return getError
	}
	if comment.Uid == uid { // Comment owner
		_, deleteError := db.c.Query(queries.RemoveComment, cid, pid)
		return deleteError
	} else {
		post, getError := db.GetPost(pid)
		if getError != nil {
			return getError
		}
		if post.Uid == uid { // Post Owner
			_, deleteError := db.c.Query(queries.RemoveComment, cid, pid)
			return deleteError
		} else {
			return errors.New("Unauthorized")
		}
	}
}
