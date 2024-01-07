package database

import (
	"crypto/rand"
	"database/sql"
	"net/http"
	"strings"

	"github.com/SimoneDiCesare/WasaPhoto/service/cypher"
	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func createUser(db *appdbimpl, username string) (user User, err error) {
	// TODO: Check if the uuid and token are really unique.
	user.Username = username
	user.Uid, err = newUserId()
	if err != nil {
		return user, err
	}
	user.Token, err = cypher.GenerateAuthToken(user.Uid)
	if err != nil {
		return user, err
	}
	_, err = db.c.Exec(queries.CreateNewUser, user.Uid, username, user.Token)
	if err != nil {
		return user, err
	}
	return user, nil
}

func newUserId() (string, error) {
	length := 12
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	charSet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var randomIDBuilder strings.Builder
	for _, b := range bytes {
		randomIDBuilder.WriteByte(charSet[int(b)%len(charSet)])
	}
	randomID := randomIDBuilder.String()
	return randomID, nil
}

func (db *appdbimpl) LoginUser(username string) (int, User, error) {
	var uid, dummyUsername, biography, token string
	scanError := db.c.QueryRow(queries.GetUserFromUsername, username).Scan(&uid, &dummyUsername, &biography, &token)
	if scanError == sql.ErrNoRows {
		// User does not exist, create a new one
		db.logger.Infof("Creating new user '%s'", username)
		user, creationError := createUser(db, username)
		return http.StatusCreated, user, creationError
	}
	user := User{
		Uid:       uid,
		Username:  username,
		Biography: biography,
		Token:     token,
	}
	// Retrieved user infos
	db.logger.Infof("'%s' loggin", username)
	return http.StatusOK, user, nil
}
