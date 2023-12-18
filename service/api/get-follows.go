package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getFollows(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	follows, retrieveingError := rt.db.GetFollows(uid)
	if retrieveingError != nil {
		rt.baseLogger.WithError(retrieveingError).Error("Error getting follows")
		http.Error(w, "Error getting follows", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content", "application/json")
	content, encodingError := json.Marshal(follows)
	if encodingError != nil {
		rt.baseLogger.WithError(encodingError).Error("Error while encoding follows")
		http.Error(w, "Error while encoding follows", http.StatusInternalServerError)
		return
	}
	_, writeError := w.Write(content)
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
