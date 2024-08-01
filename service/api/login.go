package api

import (
	"encoding/json"
	"io"
	"net/http"

	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bodyContent, readingError := io.ReadAll(r.Body)
	if readingError != nil {
		rt.baseLogger.WithError(readingError).Error("Error reading request body")
		http.Error(w, "Error reading request Body", http.StatusInternalServerError)
		return
	}
	var body schema.LoginRequestBody
	unmarshalError := json.Unmarshal(bodyContent, &body)
	if unmarshalError != nil {
		rt.baseLogger.WithError(unmarshalError).Error("Error decoding json body")
		http.Error(w, "Error decoding json body", http.StatusInternalServerError)
		return
	}
	if body.Username == "" {
		rt.baseLogger.Error("Invalid request body: Empty")
		http.Error(w, "Invalid request body: Empty", http.StatusBadRequest)
		return
	}
	var user *schema.UserLogin
	user, dbError := rt.db.LoginUser(body.Username)
	if dbError != nil {
		rt.baseLogger.WithError(dbError).Error("Error loggin-in")
		http.Error(w, "Error loggin-in", http.StatusInternalServerError)
		return
	}
	returnData, marshalError := json.Marshal(user)
	if marshalError != nil {
		rt.baseLogger.WithError(marshalError).Error("Error encoding response")
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	_, writeError := w.Write(returnData)
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error writing response")
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
