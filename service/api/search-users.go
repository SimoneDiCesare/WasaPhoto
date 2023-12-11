package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) SearchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	searchedText := r.URL.Query().Get("query")
	users, searchError := rt.db.SearchUsers(r.Header.Get("Token"), searchedText)
	if searchError != nil {
		rt.baseLogger.WithError(searchError).Error("Error while searching query")
		http.Error(w, "Error while searching query", http.StatusInternalServerError)
		return
	}
	content, encodingError := json.Marshal(users)
	if encodingError != nil {
		rt.baseLogger.WithError(encodingError).Error("Error while encoding response")
		http.Error(w, "Error while encoding response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, writeError := w.Write(content)
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
