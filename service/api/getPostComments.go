package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPostComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	comments, err := rt.db.GetPostComments(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error getting comments")
		http.Error(w, "Error getting comments", http.StatusInternalServerError)
		return
	}
	content, encodingError := json.Marshal(comments)
	if encodingError != nil {
		rt.baseLogger.WithError(err).Error("Error encoding response")
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, writeError := w.Write(content)
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
	}
}
