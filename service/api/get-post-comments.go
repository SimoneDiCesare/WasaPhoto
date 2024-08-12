package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPostComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reqUid, tokenError := rt.checkToken(r.Header.Get("token"))
	if tokenError != nil {
		rt.HandleTokenError(tokenError, w)
		return
	}
	pid := ps.ByName("pid")
	simplePost, simpleError := rt.db.GetSimplePost(pid)
	if simpleError != nil {
		rt.baseLogger.WithError(simpleError).Debug("Can't check bans for liking")
	} else {
		banError := rt.checkBanned(simplePost.Author.Uid, reqUid)
		if banError != nil {
			rt.HandleBanError(banError, w)
			return
		}
	}
	comments, err := rt.db.GetPostComments(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error getting Comments")
		http.Error(w, "Error getting Comments", http.StatusInternalServerError)
		return
	}
	data, marshalError := json.Marshal(comments)
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
