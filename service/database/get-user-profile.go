package database

import (
	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
)

func (db *appdbimpl) GetUserProfile(uid string) (profile *schema.UserProfileData, err error) {
	profile = &schema.UserProfileData{}
	posts, err := db.GetUserPosts(uid)
	if err != nil {
		return nil, err
	}
	profile.Posts = posts
	err = db.c.QueryRow(GetSimpleUserFromId, uid).Scan(&profile.User.Uid, &profile.User.Username)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// TOOD: Change to complete post return type
func (db *appdbimpl) GetUserPost(uid string, pid string) (*schema.SimplePostData, error) {
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

func (db *appdbimpl) GetUserPosts(uid string) (posts []schema.SimplePostData, err error) {
	rows, searchError := db.c.Query(GetUserPosts, uid)
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
		var post schema.SimplePostData
		if err := rows.Scan(&post.Pid, &post.Author.Uid, &post.Author.Username, &post.CreatedAt); err != nil {
			db.logger.Debug("Scan Error")
			return nil, err
		}
		post.ImageUrl = "users/" + post.Author.Uid + "/posts/" + post.Pid
		posts = append(posts, post)
	}
	rowsError := rows.Err()
	if rowsError != nil {
		db.logger.Debug("Rows Error")
		return nil, rowsError
	}
	return posts, nil
}
