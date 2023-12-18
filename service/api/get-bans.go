package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getBans(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid1")
	bans, banError := rt.db.GetBans(uid)
	if banError != nil {
		rt.baseLogger.WithError(banError).Error("Error while banning user")
		http.Error(w, "Error while banning user", http.StatusInternalServerError)
		return
	}
	content, encodingError := json.Marshal(bans)
	if encodingError != nil {
		rt.baseLogger.WithError(encodingError).Error("Error while encoding bans")
		http.Error(w, "Error while encoding bans", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	_, writeError := w.Write(content)
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
	}
}
