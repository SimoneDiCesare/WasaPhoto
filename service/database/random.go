package database

import (
	"crypto/rand"
	"encoding/hex"
)

func generateRandomHex() (string, error) {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	hex := hex.EncodeToString(bytes)
	return hex, nil
}
