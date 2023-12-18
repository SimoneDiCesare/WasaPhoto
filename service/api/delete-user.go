package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	deleteError := rt.db.DeleteUser(uid)
	if deleteError != nil {
		rt.baseLogger.WithError(deleteError).Error("Error deleting profile")
		http.Error(w, "Error deleting profile", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	_, writeError := w.Write([]byte("User deleted."))
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
