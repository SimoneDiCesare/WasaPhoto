package database

import (
	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
)

func (db *appdbimpl) GetSimplePost(pid string) (*schema.SimplePostData, error) {
	simplePost := &schema.SimplePostData{
		Pid: pid,
	}
	// Retrieve correct infos
	row := db.c.QueryRow(GetSimplePost, pid)
	err := row.Scan(&simplePost.Pid, &simplePost.Author.Uid,
		&simplePost.Author.Username, &simplePost.CreatedAt)
	simplePost.ImageUrl = "users/" + simplePost.Author.Uid + "/posts/" + simplePost.Pid
	if err != nil {
		db.logger.Error(err)
	}
	return simplePost, nil
}
