package api

import (
	"encoding/json"
	"net/http"

	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reqUid, tokenError := rt.checkToken(r.Header.Get("token"))
	if tokenError != nil {
		rt.HandleTokenError(tokenError, w)
		return
	}
	uid := ps.ByName("uid")
	if reqUid != uid {
		rt.baseLogger.Error("Can't like for another user!")
		http.Error(w, "Can't like for another user!", http.StatusForbidden)
		return
	}
	pid := ps.ByName("pid")
	err := rt.db.LikePost(uid, pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Can't like post")
		http.Error(w, "Can't like post", http.StatusInternalServerError)
		return
	}
	like := &schema.LikeData{}
	like.Uid = uid
	like.Pid = pid
	data, marshalError := json.Marshal(like)
	if marshalError != nil {
		rt.baseLogger.WithError(marshalError).Error("Error encoding response")
		http.Error(w, "Error Encoding Response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeError := w.Write(data)
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error writing response")
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}

func (rt *_router) unlikePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reqUid, tokenError := rt.checkToken(r.Header.Get("token"))
	if tokenError != nil {
		rt.HandleTokenError(tokenError, w)
		return
	}
	uid := ps.ByName("uid")
	if reqUid != uid {
		rt.baseLogger.Error("Can't unlike for another user!")
		http.Error(w, "Can't unlike for another user!", http.StatusForbidden)
		return
	}
	pid := ps.ByName("pid")
	err := rt.db.UnlikePost(uid, pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Can't unlike post")
		http.Error(w, "Can't unlike post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
