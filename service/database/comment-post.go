package database

import (
	"crypto/rand"
	"strings"

	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func newCommentId() (string, error) {
	length := 15
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	charSet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var randomIDBuilder strings.Builder
	for _, b := range bytes {
		randomIDBuilder.WriteByte(charSet[int(b)%len(charSet)])
	}
	randomID := randomIDBuilder.String()
	return randomID, nil
}

func (db *appdbimpl) CommentPost(comment SimpleComment) error {
	cid, cidError := newCommentId()
	if cidError != nil {
		return cidError
	}
	_, queryErr := db.c.Exec(queries.CreateNewComment, cid, comment.Pid, comment.Uid, comment.Content)
	if queryErr != nil {
		return queryErr
	}
	return nil
}
