package api

type AuthenticationError struct{}

func (m *AuthenticationError) Error() string {
	return "Not Authenticated"
}

type AuthorizationError struct{}

func (m *AuthorizationError) Error() string {
	return "Not Authorized"
}

// Check if its a valid token -> generic operations like search user
// returns the uid associated with the token
func (rt *_router) checkToken(token string) (string, error) {
	rt.baseLogger.Debugf("Token to verify: %s", token)
	uid, searchError := rt.db.SearchUidByToken(token)
	if searchError != nil {
		return "", &AuthenticationError{}
	}
	return uid, nil
}

func (rt *_router) checkTokenForUid(token string, uid string) error {
	resultUid, searchError := rt.db.SearchUidByToken(token)
	if searchError != nil {
		return &AuthenticationError{}
	}
	if resultUid != uid {
		return &AuthorizationError{}
	}
	return nil
}
