package database

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
)

func (db *appdbimpl) CreatePost(uid string) (string, error) {
	triesCount := 0
	var pid string
	var foundId string
	for {
		pid = fmt.Sprintf("%x", rand.Uint64())
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
		return "", nil
	}
	return pid, nil
}
