package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reqUid, tokenError := rt.checkToken(r.Header.Get("Authorization"))
	if tokenError != nil {
		rt.HandleTokenError(tokenError, w)
		return
	}
	uid := ps.ByName("uid")
	banError := rt.checkBanned(uid, reqUid)
	if banError != nil {
		rt.HandleBanError(banError, w)
		return
	}
	rt.baseLogger.Debugf("Getting profile of %s", uid)
	profile, getError := rt.db.GetUserProfile(uid)
	if getError != nil {
		rt.baseLogger.WithError(getError).Error("Error getting Profile")
		http.Error(w, "Error getting Profile", http.StatusInternalServerError)
		return
	}
	data, marshalError := json.Marshal(profile)
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
