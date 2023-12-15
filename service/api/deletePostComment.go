package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePostComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token := r.Header.Get("Token")
	pid, cid := ps.ByName("pid"), ps.ByName("cid")
	err := rt.db.DeletePostComment(token, pid, cid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error getting comment")
		http.Error(w, "Error getting comment", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	_, writeError := w.Write([]byte("Comment deleted!"))
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
	}
}
