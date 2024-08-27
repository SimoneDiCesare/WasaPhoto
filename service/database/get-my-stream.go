package database

import (
	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
)

func (db *appdbimpl) GetMyStream(uid string) (posts []schema.SimplePostData, err error) {
	rows, searchError := db.c.Query(GetFeeds, uid)
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
