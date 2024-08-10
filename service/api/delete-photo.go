package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse multipart body -> 20 MB
	uid, tokenErr := rt.checkToken(r.Header.Get("token"))
	if tokenErr != nil {
		rt.HandleTokenError(tokenErr, w)
		return
	}
	pid := ps.ByName("pid")
	err := rt.db.DeletePhoto(uid, pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error deleting photo")
		http.Error(w, "Error deleting photo", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
