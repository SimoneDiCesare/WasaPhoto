package database

import (
	"errors"

	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func (db *appdbimpl) GetFollows(uid string) ([]SimpleUserProfile, error) {
	var followers []SimpleUserProfile
	rows, getFollowersError := db.c.Query(queries.GetFollows, uid)
	if getFollowersError != nil {
		return nil, errors.New("Error retrieving followers")
	}
	defer rows.Close()
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
