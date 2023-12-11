package database

import (
	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func (db *appdbimpl) SearchUsers(token string, text string) (users []SimpleUserProfile, err error) {
	var uid string
	uidError := db.c.QueryRow(queries.GetUseridFromToken, token).Scan(&uid)
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
	return users, nil
}
