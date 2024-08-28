package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid, tokenError := rt.checkToken(r.Header.Get("Authorization"))
	if tokenError != nil {
		rt.baseLogger.WithError(tokenError).Error("Unauthorized")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	textToSearch := r.URL.Query().Get("username")
	if textToSearch == "" {
		rt.baseLogger.Error("Empty search on username")
		http.Error(w, "Empty search on username", http.StatusBadRequest)
		return
	}
	users, searchError := rt.db.SearchUsersByName(uid, textToSearch)
	if searchError != nil {
		rt.baseLogger.WithError(searchError).Error("Error searching users")
		http.Error(w, "Error searching users", http.StatusInternalServerError)
		return
	}
	data, marshalError := json.Marshal(users)
	if marshalError != nil {
		rt.baseLogger.WithError(marshalError).Error("Error encoding response")
		http.Error(w, "Error Encoding Response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	_, writeError := w.Write(data)
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error writing response")
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
