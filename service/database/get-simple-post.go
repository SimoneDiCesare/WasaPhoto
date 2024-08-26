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
<<<<<<< HEAD

	simplePost.ImageUrl = "/posts/" + simplePost.Pid + "/image.png"
=======
	simplePost.ImageUrl = "users/" + simplePost.Author.Uid + "/posts/" + simplePost.Pid + ".png"
>>>>>>> refs/remotes/origin/main
	if err != nil {
		db.logger.Error(err)
	}
	return simplePost, nil
}
