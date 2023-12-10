package database

import "github.com/SimoneDiCesare/WasaPhoto/service/database/queries"

func (db *appdbimpl) GetPost(pid string) (*Post, error) {
	var post Post
	getPostError := db.c.QueryRow(queries.GetBasePost, pid).Scan(&post.Pid, &post.Uid, &post.Caption, &post.CreatedAt)
	if getPostError != nil {
		return nil, getPostError
	}
	post.Image = "/posts/" + post.Pid + "/image"
	return nil, nil
}
