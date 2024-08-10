package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reqUid, tokenError := rt.checkToken(r.Header.Get("token"))
	if tokenError != nil {
		rt.HandleTokenError(tokenError, w)
		return
	}
	uid := ps.ByName("uid")
	pid := ps.ByName("pid")
	banError := rt.checkBanned(uid, reqUid)
	if banError != nil {
		rt.HandleBanError(banError, w)
		return
	}
	rt.baseLogger.Debugf("Getting post of %s, %s", uid, pid)
	post, getError := rt.db.GetUserPost(uid, pid)
	if getError != nil {
		rt.baseLogger.WithError(getError).Error("Error getting Post")
		http.Error(w, "Error getting Post", http.StatusInternalServerError)
		return
	}
	data, marshalError := json.Marshal(post)
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
