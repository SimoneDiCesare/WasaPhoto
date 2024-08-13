package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	bodyContent, readError := io.ReadAll(r.Body)
	if readError != nil {
		rt.baseLogger.WithError(readError).Error("Error reading body")
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}
	content := string(bodyContent[:])
	comment, err := rt.db.CommentPhoto(pid, reqUid, content)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Can't comment post")
		http.Error(w, "CAn't comment post", http.StatusInternalServerError)
		return
	}
	data, marshalError := json.Marshal(comment)
	if marshalError != nil {
		rt.baseLogger.WithError(marshalError).Error("Error encoding response")
		http.Error(w, "Error Encoding Response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, writeError := w.Write(data)
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error writing response")
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	cid := ps.ByName("cid")
	err := rt.db.UncommentPhoto(cid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Can't delete post comment")
		http.Error(w, "Can't delete post comment", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
