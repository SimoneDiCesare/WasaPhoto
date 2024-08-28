package database

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"

	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
)

func (db *appdbimpl) CommentPhoto(pid string, uid string, content string) (*schema.PostComment, error) {
	triesCount := 0
	var cid string
	var foundId string
	for {
		cid = fmt.Sprintf("%x", rand.Uint64())
		queryError := db.c.QueryRow(GetCommentIdFromId, cid).Scan(&foundId)
		if errors.Is(queryError, sql.ErrNoRows) {
			break
		}
		triesCount = triesCount + 1
		if triesCount > 10 {
			return nil, errors.New("exceeded tries for creating cid")
		}
	}
	_, err := db.c.Exec(CommentPhoto, cid, pid, uid, content)
	if err != nil {
		db.logger.Debugf("Error commenting: %e", err)
		return nil, err
	}
	var user schema.SimpleUserData
	err = db.c.QueryRow(GetSimpleUserFromId, uid).Scan(&user.Uid, &user.Username)
	if err != nil {
		db.logger.Debugf("Can't retrieve simpleuser: %e", err)
		return nil, err
	}
	comment := &schema.PostComment{
		Pid:    pid,
		Author: user,
		Cid:    cid,
		Text:   content,
	}
	return comment, nil
}

func (db *appdbimpl) UncommentPhoto(cid string) error {
	_, err := db.c.Exec(UncommentPhoto, cid)
	if err != nil {
		db.logger.Debugf("Error uncomment post: %e", err)
		return err
	}
	return nil
}
