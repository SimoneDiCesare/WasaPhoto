package api

import (
	"encoding/json"
	"net/http"

	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ownerId, tokenError := rt.checkToken(r.Header.Get("token"))
	if tokenError != nil {
		rt.HandleTokenError(tokenError, w)
		return
	}
	// Only the proprietary can ban users
	uid := ps.ByName("uid")
	if ownerId != uid {
		rt.baseLogger.Error("Forbidden Ban")
		http.Error(w, "Forbidden ban", http.StatusForbidden)
		return
	}
	bid := ps.ByName("bid")
	// Can't ban someone who banned me first
	banError := rt.checkBanned(bid, ownerId)
	if banError != nil {
		rt.HandleBanError(banError, w)
		return
	}
	if uid == bid {
		rt.baseLogger.Error("Can't ban itself")
		http.Error(w, "Can't ban itself", http.StatusBadRequest)
		return
	}
	bans, getError := rt.db.GetBans(uid)
	if getError != nil {
		rt.baseLogger.WithError(getError).Error("Error getting Bans")
		http.Error(w, "Error getting Bans", http.StatusInternalServerError)
		return
	}
	data, marshalError := json.Marshal(schema.BanData{
		BanningId: uid,
		BannedId:  bid,
	})
	if marshalError != nil {
		rt.baseLogger.WithError(marshalError).Error("Error encoding response")
		http.Error(w, "Error Encoding Response", http.StatusInternalServerError)
		return
	}
	for _, ban := range bans {
		if ban.Uid == bid {
			// Already bannes
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
	// Create ban and return
	banError = rt.db.BanUser(uid, bid)
	if banError != nil {
		rt.baseLogger.WithError(banError).Error("Error Banning user")
		http.Error(w, "Error Banning user", http.StatusInternalServerError)
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

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ownerId, tokenError := rt.checkToken(r.Header.Get("token"))
	if tokenError != nil {
		rt.HandleTokenError(tokenError, w)
		return
	}
	uid := ps.ByName("uid")
	if ownerId != uid {
		rt.baseLogger.Error("Forbidden unban")
		http.Error(w, "Forbidden unban", http.StatusForbidden)
		return
	}
	bid := ps.ByName("bid")
	// Can't unban someone who banned me first
	banError := rt.checkBanned(bid, ownerId)
	if banError != nil {
		rt.HandleBanError(banError, w)
		return
	}
	if uid == bid {
		rt.baseLogger.Error("Can't unban itself")
		http.Error(w, "Can't unban itself", http.StatusBadRequest)
		return
	}
	dbError := rt.db.UnbanUser(uid, bid)
	if dbError != nil {
		rt.baseLogger.WithError(dbError).Error("Error unban user")
		http.Error(w, "Error unban user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
