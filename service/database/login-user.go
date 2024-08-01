package database

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"

	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
)

func (db *appdbimpl) CreateUser(username string) (*schema.UserLogin, error) {
	var user schema.UserLogin
	user.Username = username
	// Generate a valid id for the new user
	// If on the 10-th tries we can't generate a valid id, an error is returned.
	triesCount := 0
	for {
		user.Uid = fmt.Sprintf("%x", rand.Uint64())
		queryError := db.c.QueryRow(GetUserById, user.Uid).Scan()
		if queryError == sql.ErrNoRows {
			break
		}
		triesCount++
		if triesCount > 10 {
			return nil, errors.New("exceeded tries for creating uid")
		}
	}
	user.Token = fmt.Sprintf("%x", rand.Uint64())
	_, queryError := db.c.Exec(CreateUser, user.Uid, user.Username, user.Token)
	if queryError != nil {
		return nil, queryError
	}
	return &user, nil
}

func (db *appdbimpl) LoginUser(username string) (*schema.UserLogin, error) {
	var user schema.UserLogin
	queryError := db.c.QueryRow(GetUserByName, username).Scan(&user.Uid, &user.Username, &user.Token)
	if queryError == sql.ErrNoRows {
		db.logger.Info("Creating new user: " + username)
		return db.CreateUser(username)
	}
	// Update token on login
	tmpToken := fmt.Sprintf("%x", rand.Uint64())
	_, queryError = db.c.Exec(UpdateUserToken, user.Uid, tmpToken)
	if queryError != nil {
		// Use old token
		db.logger.Info("Can't update older token")
		db.logger.Info(queryError)
		return &user, nil
	}
	user.Token = tmpToken
	db.logger.Info("User logged in")
	return &user, nil
}
