package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	post, getError := rt.db.GetPost(pid)
	if getError != nil {
		rt.baseLogger.WithError(getError).Error("Error getting post")
		http.Error(w, "Error geting post", http.StatusInternalServerError)
		return
	}
	content, encodingError := json.Marshal(post)
	if encodingError != nil {
		rt.baseLogger.WithError(encodingError).Error("Error encoding response")
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, writeError := w.Write(content)
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error writing response")
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
