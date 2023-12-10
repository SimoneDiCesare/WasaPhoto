package database

import (
	"crypto/rand"
	"errors"
	"strings"

	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func newPostId() (string, error) {
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

func (db *appdbimpl) CreatePost(uid string, caption string) (*Post, error) {
	pid, pidError := newPostId()
	if pidError != nil {
		return nil, pidError
	}
	_, createError := db.c.Exec(queries.CreateNewPost, pid, uid, caption)
	if createError != nil {
		return nil, createError
	}
	var post Post
	getPostError := db.c.QueryRow(queries.GetBasePost, pid).Scan(&post.Pid, &post.Uid, &post.Caption, &post.CreatedAt)
	if getPostError != nil {
		return nil, getPostError
	}
	if pid != post.Pid || uid != post.Uid {
		return nil, errors.New("Inconsistency of data")
	}
	return &post, nil
}
