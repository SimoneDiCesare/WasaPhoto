package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type SessionRequestBody struct {
	Username string `json:"username"`
}

/*
users: {uid|username|biography|token}

	if (exists(user with username)) {
		return user
	} else {

		generateUid while uid is unique
		generate Token while token is unique
		return new user(uid, username, "", token)
	}
*/
func (rt *_router) postUserSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rt.baseLogger.Infof("Login User")
	bodyContent, readingError := io.ReadAll(r.Body)
	if readingError != nil {
		rt.baseLogger.WithError(readingError).Error("Error reading request body")
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	var body SessionRequestBody
	decodingError := json.Unmarshal(bodyContent, &body)
	if decodingError != nil {
		rt.baseLogger.WithError(decodingError).Error("Error parsing request body")
		http.Error(w, "Error parsing request body", http.StatusInternalServerError)
		return
	}
	if body.Username == "" {
		rt.baseLogger.Error("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	statusCode, user, loginError := rt.db.LoginUser(body.Username)
	if loginError != nil {
		rt.baseLogger.WithError(loginError).Error("Error login response")
		http.Error(w, "Error login response", http.StatusInternalServerError)
		return
	}
	content, encodingError := json.Marshal(user)
	if encodingError != nil {
		rt.baseLogger.WithError(encodingError).Error("Error encoding response")
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	_, writingError := w.Write(content)
	if writingError != nil {
		rt.baseLogger.WithError(writingError).Error("Error while writing response")
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
