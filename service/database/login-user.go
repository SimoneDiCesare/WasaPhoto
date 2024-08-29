package database

import (
	"database/sql"
	"errors"

	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
)

// Print all Users
// rows, _ := db.c.Query(GetUsers)
// for rows.Next() {
// 	var user schema.UserLogin
// 	if err := rows.Scan(&user.Uid, &user.Username, &user.Token); err != nil {
// 		db.logger.Error("Scan Error")
// 		continue
// 	}
// 	db.logger.Infof("User: {%s, %s, %s}", user.Uid, user.Username, user.Token)
// }

func (db *appdbimpl) CreateUser(username string) (*schema.UserLogin, error) {
	var user schema.UserLogin
	user.Username = username
	var tmp1, tmp2, tmp3 string
	queryError := db.c.QueryRow(GetUserByName, username).Scan(&tmp1, &tmp2, &tmp3)
	if errors.Is(queryError, nil) {
		// Username already existing
		return nil, schema.ErrExistingUsername
	}
	// Generate a valid id for the new user
	// If on the 10-th tries we can't generate a valid id, an error is returned.
	triesCount := 0
	for {
		tmpUid, err := generateRandomHex()
		if err != nil {
			db.logger.Errorf("Error generating hex: %e", err)
			triesCount = triesCount + 1
			continue
		}
		user.Uid = tmpUid
		queryError = db.c.QueryRow(GetUserById, user.Uid).Scan()
		if errors.Is(queryError, sql.ErrNoRows) {
			break
		}
		triesCount++
		if triesCount > 10 {
			return nil, errors.New("exceeded tries for creating uid")
		}
	}
	tmpToken, err := generateRandomHex()
	if err != nil {
		return nil, err
	}
	user.Token = tmpToken
	_, queryError = db.c.Exec(CreateUser, user.Uid, user.Username, user.Token)
	if queryError != nil {
		return nil, queryError
	}
	return &user, nil
}

func (db *appdbimpl) LoginUser(username string) (*schema.UserLogin, error) {
	var user schema.UserLogin
	queryError := db.c.QueryRow(GetUserByName, username).Scan(&user.Uid, &user.Username, &user.Token)
	if errors.Is(queryError, sql.ErrNoRows) {
		db.logger.Debug("Creating new user: " + username)
		return db.CreateUser(username)
	}
	// Update token on login
	tmpToken, err := generateRandomHex()
	if err != nil {
		db.logger.Debugf("Can't update older token: %e", err)
		return &user, nil
	}
	res, queryError := db.c.Exec(UpdateUserToken, tmpToken, user.Uid)
	if queryError != nil {
		db.logger.Debugf("Can't update older token: %e", queryError)
		return &user, nil
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		// Use old token
		db.logger.Debugf("Can't update older token: %d, %e", affected, queryError)
		return &user, nil
	}
	user.Token = tmpToken
	db.logger.Debug("User logged in")
	return &user, nil
}
