package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid, uid2 := ps.ByName("uid1"), ps.ByName("uid2")
	banError := rt.db.BanUser(uid, uid2)
	if banError != nil {
		rt.baseLogger.WithError(banError).Error("Error while banning user")
		http.Error(w, "Error while banning user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	_, writeError := w.Write([]byte("User banned!"))
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
	}
}
