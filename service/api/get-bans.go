package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getBans(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	tokenError := rt.checkTokenForUid(r.Header.Get("Authorization"), uid)
	if tokenError != nil {
		rt.HandleTokenError(tokenError, w)
		return
	}
	bans, getError := rt.db.GetBans(uid)
	if getError != nil {
		rt.baseLogger.WithError(getError).Error("Error getting Bans")
		http.Error(w, "Error getting Bans", http.StatusInternalServerError)
		return
	}
	data, marshalError := json.Marshal(bans)
	if marshalError != nil {
		rt.baseLogger.WithError(marshalError).Error("Error encoding response")
		http.Error(w, "Error Encoding Response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	_, writeError := w.Write(data)
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error writing response")
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
