package database

import (
	"database/sql"
	"errors"

	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func (db *appdbimpl) GetFollowerCount(uid string) (int, error) {
	var count int
	err := db.c.QueryRow(queries.GetFollowerCount, uid).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *appdbimpl) GetFollowsCount(uid string) (int, error) {
	var count int
	err := db.c.QueryRow(queries.GetFollowsCount, uid).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *appdbimpl) GetUserProfile(uid string) (_ *UserProfile, err error) {
	var user User
	getUserError := db.c.QueryRow(queries.GetUserFromUid, uid).Scan(&user.Uid, &user.Username, &user.Biography, &user.Token)
	if getUserError == sql.ErrNoRows {
		return nil, errors.New("No user found")
	}
	var posts []Post
	rows, getUserPostsError := db.c.Query(queries.GetUserPosts, uid)
	if getUserPostsError != nil {
		return nil, errors.New("Error retrieving posts")
	}
	defer func() {
		closeErr := rows.Close()
		if err == nil {
			err = closeErr
		}
	}()
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.Pid, &post.Uid, &post.Caption, &post.CreatedAt); err != nil {
			return nil, err
		}
		post.Image = "/posts/" + post.Pid + "/image"
		posts = append(posts, post)
	}
	rowsErr := rows.Err()
	if rowsErr != nil {
		return nil, rowsErr
	}
	profileImage := "/users/" + uid + "/image"
	follower, getCountError := db.GetFollowerCount(uid)
	if getCountError != nil {
		return nil, getCountError
	}
	follows, getCountError := db.GetFollowsCount(uid)
	if getCountError != nil {
		return nil, getCountError
	}
	return &UserProfile{
		User:         user,
		Posts:        posts,
		ProfileImage: profileImage,
		Follower:     follower,
		Follows:      follows,
	}, nil
}
