package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	token := r.Header.Get("Token")
	likeError := rt.db.LikePost(pid, token)
	if likeError != nil {
		rt.baseLogger.WithError(likeError).Error("Error while liking post")
		http.Error(w, "Error while liking post", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	_, writeError := w.Write([]byte("Post Liked!"))
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
