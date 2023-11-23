package cypher

import (
	"encoding/base64"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/SimoneDiCesare/WasaPhoto/service/database"
	"github.com/golang-jwt/jwt/v5"
)

var rng *rand.Rand

func InitRng() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GenerateAuthToken(user database.User) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)
	key, err := strconv.ParseInt(strings.ToLower(user.Uid), 36, 64)
	key = key * rng.Int63()
	finalKey, err := base64.StdEncoding.DecodeString(strconv.FormatInt(int64(key), 10))
	s, err := t.SignedString(finalKey)
	if err != nil {
		return "", err
	}
	return s, nil
}
