package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deleteError := rt.db.DeletePost(ps.ByName("pid"), r.Header.Get("Token"))
	if deleteError != nil {
		rt.baseLogger.WithError(deleteError).Error("Error while deleting post")
		http.Error(w, "Error while deleting post", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	_, writeError := w.Write([]byte("Post deleted!"))
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
	}
}
