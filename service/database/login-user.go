package database

import (
	"crypto/rand"
	"database/sql"
	"strings"

	"github.com/SimoneDiCesare/WasaPhoto/service/cypher"
	"github.com/SimoneDiCesare/WasaPhoto/service/database/queries"
)

func createUser(db *appdbimpl, username string) (uid string, token string, err error) {
	// TODO: Check if the uuid and token are really unique.
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
	charSet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var randomIDBuilder strings.Builder
	for _, b := range bytes {
		randomIDBuilder.WriteByte(charSet[int(b)%len(charSet)])
	}
	randomID := randomIDBuilder.String()
	return randomID, nil
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
