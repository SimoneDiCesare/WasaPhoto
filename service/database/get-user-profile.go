package database

import (
	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
)

func (db *appdbimpl) GetUserProfile(uid string) (profile *schema.UserProfileData, err error) {
	profile = &schema.UserProfileData{}
	posts, err := db.GetUserPosts(uid)
	if err != nil {
		db.logger.Debug("Error on user posts")
		return nil, err
	}
	profile.Posts = posts
	err = db.c.QueryRow(GetSimpleUserFromId, uid).Scan(&profile.User.Uid, &profile.User.Username)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (db *appdbimpl) GetPostLikes(pid string) (likes []schema.SimpleUserData, err error) {
	rows, searchError := db.c.Query(GetPostLikes, pid)
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
		var like schema.SimpleUserData
		if err := rows.Scan(&like.Uid, &like.Username); err != nil {
			db.logger.Debug("Scan Error")
			return nil, err
		}
		likes = append(likes, like)
	}
	rowsError := rows.Err()
	if rowsError != nil {
		db.logger.Debug("Rows Error")
		return nil, rowsError
	}
	return likes, nil
}

func (db *appdbimpl) GetUserPost(uid string, pid string) (*schema.PostData, error) {
	post := &schema.PostData{
		Pid: pid,
	}
	// Get generic infos
	row := db.c.QueryRow(GetSimplePost, pid)
	err := row.Scan(&post.Pid, &post.Author.Uid,
		&post.Author.Username, &post.UploadTime)
	if err != nil {
		db.logger.Error(err)
		return nil, err
	}
	post.ImageUrl = "http://" + db.host + "/posts/" + post.Pid + "/image.png"
	// Get comments
	post.Comments, err = db.GetPostComments(pid)
	if err != nil {
		db.logger.Error(err)
		return nil, err
	}
	post.CommentsCount = len(post.Comments)
	// Get likes
	post.Likes, err = db.GetPostLikes(pid)
	if err != nil {
		db.logger.Error(err)
		return nil, err
	}
	post.LikesCount = len(post.Likes)
	return post, nil
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
		post.ImageUrl = "http://" + db.host + "/posts/" + post.Pid + "/image.png"
		posts = append(posts, post)
	}
	rowsError := rows.Err()
	if rowsError != nil {
		db.logger.Debug("Rows Error")
		return nil, rowsError
	}
	return posts, nil
}
