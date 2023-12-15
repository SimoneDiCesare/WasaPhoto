package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unlikePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	token := r.Header.Get("Token")
	unlikeError := rt.db.UnlikePost(pid, token)
	if unlikeError != nil {
		rt.baseLogger.WithError(unlikeError).Error("Error while unliking post")
		http.Error(w, "Error while unliking post", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/plain")
	_, writeError := w.Write([]byte("Post Unliked!"))
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
