package database

import (
	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func (db *appdbimpl) SearchUsers(token string, text string) (users []SimpleUserProfile, err error) {
	uid, uidError := db.GetUserIdFromToken(token)
	if uidError != nil {
		return nil, uidError
	}
	rows, searchUsersError := db.c.Query(queries.SearchUser, uid, text)
	if searchUsersError != nil {
		return nil, searchUsersError
	}
	defer func() {
		closeErr := rows.Close()
		if err == nil {
			err = closeErr
		}
	}()
	for rows.Next() {
		var user SimpleUserProfile
		if err := rows.Scan(&user.Uid, &user.ProfileImage); err != nil {
			return nil, err
		}
		user.ProfileImage = "/users/" + user.Uid + "/image"
		users = append(users, user)
	}
	rowsErr := rows.Err()
	if rowsErr != nil {
		return nil, rowsErr
	}
	return users, nil
}
