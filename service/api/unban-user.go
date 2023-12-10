package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid, uid2 := ps.ByName("uid1"), ps.ByName("uid2")
	banError := rt.db.UnbanUser(uid, uid2)
	if banError != nil {
		rt.baseLogger.WithError(banError).Error("Error while unbanning user")
		http.Error(w, "Error while unbanning user", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/plain")
	_, writeError := w.Write([]byte("User unbanned!"))
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
	}
}
