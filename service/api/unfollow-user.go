package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid, uid2 := ps.ByName("uid1"), ps.ByName("uid2")
	followError := rt.db.UnfollowUser(uid, uid2)
	if followError != nil {
		rt.baseLogger.WithError(followError).Error("Error while following user")
		http.Error(w, "Error while following user", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, writeError := w.Write([]byte("User unfollowed!"))
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
