package database

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"

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
	if queryError == nil {
		// Username already existing
		return nil, schema.ErrExistingUsername
	}
	// Generate a valid id for the new user
	// If on the 10-th tries we can't generate a valid id, an error is returned.
	triesCount := 0
	for {
		user.Uid = fmt.Sprintf("%x", rand.Uint64())
		queryError = db.c.QueryRow(GetUserById, user.Uid).Scan()
		if errors.Is(queryError, sql.ErrNoRows) {
			break
		}
		triesCount++
		if triesCount > 10 {
			return nil, errors.New("exceeded tries for creating uid")
		}
	}
	user.Token = fmt.Sprintf("%x", rand.Uint64())
	_, queryError = db.c.Exec(CreateUser, user.Uid, user.Username, user.Token)
	if queryError != nil {
		return nil, queryError
	}
	return &user, nil
}

func (db *appdbimpl) LoginUser(username string) (*schema.UserLogin, error) {
	var user schema.UserLogin
	queryError := db.c.QueryRow(GetUserByName, username).Scan(&user.Uid, &user.Username, &user.Token)
	if queryError == sql.ErrNoRows {
		db.logger.Debug("Creating new user: " + username)
		return db.CreateUser(username)
	}
	// Update token on login
	tmpToken := fmt.Sprintf("%x", rand.Uint64())
	res, queryError := db.c.Exec(UpdateUserToken, tmpToken, user.Uid)
	affected, _ := res.RowsAffected()
	if queryError != nil || affected == 0 {
		// Use old token
		db.logger.Debugf("Can't update older token: %d, %w", affected, queryError)
		return &user, nil
	}
	user.Token = tmpToken
	db.logger.Debug("User logged in")
	return &user, nil
}
