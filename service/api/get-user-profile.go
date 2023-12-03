package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	user, userProfileError := rt.db.GetUserProfile(uid)
	if userProfileError != nil {
		rt.baseLogger.WithError(userProfileError).Error("Error getting user profile")
		http.Error(w, "Error getting user profile", http.StatusInternalServerError)
		return
	}
	content, encodeError := json.Marshal(user)
	if encodeError != nil {
		rt.baseLogger.WithError(encodeError).Error("Error encoding user")
		http.Error(w, "Error encoding user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content", "application/json")
	_, writeError := w.Write(content)
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func (rt *_router) redirectoToUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	url := "/users/" + ps.ByName("uid") + "/profile"
	rt.baseLogger.Infof("Redirecting request to %s", url)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
