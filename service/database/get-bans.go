package database

import (
	"errors"

	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func (db *appdbimpl) GetBans(uid string) ([]SimpleUserProfile, error) {
	var bans []SimpleUserProfile
	rows, getBansError := db.c.Query(queries.GetBans, uid)
	if getBansError != nil {
		return nil, errors.New("Error retrieving bans")
	}
	defer rows.Close()
	for rows.Next() {
		var ban SimpleUserProfile
		if err := rows.Scan(&ban.Uid, &ban.Username); err != nil {
			return nil, err
		}
		ban.ProfileImage = "/users/" + ban.Uid + "/image.png"
		bans = append(bans, ban)
	}
	return bans, nil
}
