package database

import (
	"errors"

	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func (db *appdbimpl) GetFollowers(uid string) (followers []SimpleUserProfile, err error) {
	rows, getFollowersError := db.c.Query(queries.GetFollowers, uid)
	if getFollowersError != nil {
		return nil, errors.New("Error retrieving followers")
	}
	defer func() {
		closeErr := rows.Close()
		if err == nil {
			err = closeErr
		}
	}()
	for rows.Next() {
		var follower SimpleUserProfile
		if err := rows.Scan(&follower.Uid, &follower.Username); err != nil {
			return nil, err
		}
		follower.ProfileImage = "/users/" + follower.Uid + "/image.png"
		followers = append(followers, follower)
	}
	return followers, nil
}
