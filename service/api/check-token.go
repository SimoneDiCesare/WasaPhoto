package api

import schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"

// Check if its a valid token -> generic operations like search user
// returns the uid associated with the token
func (rt *_router) checkToken(token string) (string, error) {
	rt.baseLogger.Debugf("Token to verify: %s", token)
	uid, searchError := rt.db.SearchUidByToken(token)
	if searchError != nil {
		return "", schema.ErrNoAuthentication
	}
	return uid, nil
}

func (rt *_router) checkTokenForUid(token string, uid string) error {
	resultUid, searchError := rt.db.SearchUidByToken(token)
	if searchError != nil {
		return schema.ErrNoAuthentication
	}
	if resultUid != uid {
		return schema.ErrNotAuthorized
	}
	return nil
}

func (rt *_router) checkBanned(uid1 string, uid2 string) error {
	// Same user has no ban problems
	if uid1 == uid2 {
		return nil
	}
	bans, searchError := rt.db.GetBans(uid1)
	if searchError != nil {
		return searchError
	}
	for _, ban := range bans {
		if ban.Uid == uid2 {
			return schema.ErrNotAuthorized
		}
	}
	return nil
}
