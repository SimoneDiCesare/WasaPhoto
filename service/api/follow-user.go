package api

import (
	"encoding/json"
	"net/http"

	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, tokenError := rt.checkToken(r.Header.Get("token"))
	if tokenError != nil {
		rt.HandleTokenError(tokenError, w)
		return
	}
	uid := ps.ByName("uid")
	fid := ps.ByName("fid")
	if uid == fid {
		rt.baseLogger.Error("Can't follow itself")
		http.Error(w, "Can't follow itself", http.StatusBadRequest)
		return
	}
	follows, getError := rt.db.GetFollowers(uid)
	if getError != nil {
		rt.baseLogger.WithError(getError).Error("Error getting Follows")
		http.Error(w, "Error getting Follows", http.StatusInternalServerError)
		return
	}
	data, marshalError := json.Marshal(schema.FollowData{
		FollowingId: fid,
		FollowedId:  uid,
	})
	if marshalError != nil {
		rt.baseLogger.WithError(marshalError).Error("Error encoding response")
		http.Error(w, "Error Encoding Response", http.StatusInternalServerError)
		return
	}
	for _, follow := range follows {
		if follow.Uid == fid {
			// Already followed
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, writeError := w.Write(data)
			if writeError != nil {
				rt.baseLogger.WithError(writeError).Error("Error writing response")
				http.Error(w, "Error writing response", http.StatusInternalServerError)
				return
			}
			return
		}
	}
	followError := rt.db.FollowUser(uid, fid)
	if followError != nil {
		rt.baseLogger.WithError(followError).Error("Error following user")
		http.Error(w, "Error following user", http.StatusInternalServerError)
		return
	}
	// Create follow link and return
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, writeError := w.Write(data)
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error writing response")
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, tokenError := rt.checkToken(r.Header.Get("token"))
	if tokenError != nil {
		rt.HandleTokenError(tokenError, w)
		return
	}
	uid := ps.ByName("uid")
	fid := ps.ByName("fid")
	if uid == fid {
		rt.baseLogger.Error("Can't unfollow itself")
		http.Error(w, "Can't unfollow itself", http.StatusBadRequest)
		return
	}
	dbError := rt.db.UnfollowUser(uid, fid)
	if dbError != nil {
		rt.baseLogger.WithError(dbError).Error("Error unfollowing user")
		http.Error(w, "Error unfollowing user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
