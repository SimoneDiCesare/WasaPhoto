package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPostLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	likes, getError := rt.db.GetPostLikes(pid)
	if getError != nil {
		rt.baseLogger.WithError(getError).Error("Error while retrieving likes")
		http.Error(w, "Error while retrieving likes", http.StatusInternalServerError)
		return
	}
	content, encodingError := json.Marshal(likes)
	if encodingError != nil {
		rt.baseLogger.WithError(getError).Error("Error while encoding response")
		http.Error(w, "Error while encoding response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, writeError := w.Write(content)
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
	}
}
