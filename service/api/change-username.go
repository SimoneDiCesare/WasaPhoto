package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) changeUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	tokenError := rt.checkTokenForUid(r.Header.Get("Authorization"), uid)
	if tokenError != nil {
		rt.HandleTokenError(tokenError, w)
		return
	}
	bodyContent, readingError := io.ReadAll(r.Body)
	if readingError != nil {
		rt.baseLogger.WithError(readingError).Error("Error reading body")
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}
	username := string(bodyContent[:])
	updateError := rt.db.ChangeUserName(uid, username)
	if errors.Is(updateError, schema.ErrExistingUsername) {
		rt.baseLogger.WithError(updateError).Error("Can't update username")
		http.Error(w, "Can't update username", http.StatusBadRequest)
		return
	} else if updateError != nil {
		rt.baseLogger.WithError(updateError).Error("Can't update username")
		http.Error(w, "Can't update username", http.StatusInternalServerError)
		return
	}
	var user schema.SimpleUserData
	user.Uid = uid
	user.Username = username
	data, marshalError := json.Marshal(user)
	if marshalError != nil {
		rt.baseLogger.WithError(updateError).Error("Error encoding response")
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
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
