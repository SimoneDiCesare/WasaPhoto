package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPostComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid, cid := ps.ByName("pid"), ps.ByName("cid")
	comment, err := rt.db.GetPostComment(pid, cid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error getting comment")
		http.Error(w, "Error getting comment", http.StatusInternalServerError)
		return
	}
	content, encodingError := json.Marshal(comment)
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
