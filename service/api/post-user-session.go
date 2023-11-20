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

type UserSession struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	BearerToken string `json:"token"`
}

func (rt *_router) postUserSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Getting username
	bodyContent, err := io.ReadAll(r.Body)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error reading request body")
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	var body SessionRequestBody
	err = json.Unmarshal(bodyContent, &body)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error decoding request body")
		http.Error(w, "Error decoding request body", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	if body.Username != "" {
		id, token, err := rt.db.LoginUser(body.Username)
		if err != nil {
			rt.baseLogger.WithError(err).Error("Error login response")
			http.Error(w, "Error login response", http.StatusInternalServerError)
			return
		}
		content, err := json.Marshal(UserSession{
			Id:          id,
			Username:    body.Username,
			BearerToken: token,
		})
		if err != nil {
			rt.baseLogger.WithError(err).Error("Error encoding response")
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
		w.Write(content)
	} else {
		rt.baseLogger.WithError(err).Error("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}
}
