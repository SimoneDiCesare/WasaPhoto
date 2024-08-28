package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	tokenError := rt.checkTokenForUid(r.Header.Get("Authorization"), uid)
	if tokenError != nil {
		rt.HandleTokenError(tokenError, w)
		return
	}
	posts, getError := rt.db.GetMyStream(uid)
	if getError != nil {
		rt.baseLogger.WithError(getError).Error("Error getting Stream")
		http.Error(w, "Error getting Stream", http.StatusInternalServerError)
		return
	}
	data, marshalError := json.Marshal(posts)
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
