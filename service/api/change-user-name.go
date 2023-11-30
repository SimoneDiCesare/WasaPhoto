package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) changeUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	body, readBodyError := io.ReadAll(r.Body)
	if readBodyError != nil {
		rt.baseLogger.WithError(readBodyError).Error("Error reading body")
		http.Error(w, "Error reading body", http.StatusInternalServerError)
	}
	var newUserName SessionRequestBody
	decodingError := json.Unmarshal(body, &newUserName)
	if decodingError != nil {
		rt.baseLogger.WithError(decodingError).Error("Error decoding request body")
		http.Error(w, "Error decoding request body", http.StatusInternalServerError)
	}
	updateError := rt.db.ChangeUserName(newUserName.Username, uid)
	if updateError != nil {
		rt.baseLogger.WithError(updateError).Error("Error updating username")
		http.Error(w, "Error updating username", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, writeError := w.Write([]byte("Username updated!"))
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error writing response")
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
