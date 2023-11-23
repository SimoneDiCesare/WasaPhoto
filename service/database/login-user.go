package database

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"

	"github.com/SimoneDiCesare/WasaPhoto/service/cypher"
	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func createUser(db *appdbimpl, username string) (uid string, token string, err error) {
	uid, err = newUserId()
	if err != nil {
		return "", "", err
	}
	token, err = cypher.GenerateAuthToken(uid)
	if err != nil {
		return "", "", err
	}
	_, err = db.c.Exec(queries.CreateNewUser, uid, username, token)
	if err != nil {
		return "", "", err
	}
	return uid, token, nil
}

func newUserId() (string, error) {
	length := 12
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	randomID := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(bytes)
	return randomID[:length], nil
}

func (db *appdbimpl) LoginUser(username string) (id string, token string, err error) {
	err = db.c.QueryRow("SELECT id, token FROM users WHERE username=$1", username).Scan(&id, &token)
	if err == sql.ErrNoRows {
		// User does not exist, create a new one
		return createUser(db, username)
	}
	// Retrieved user infos
	return id, token, err
}
