package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/SimoneDiCesare/WasaPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bodyContent, readingError := io.ReadAll(r.Body)
	if readingError != nil {
		rt.baseLogger.WithError(readingError).Error("Error while reading body")
		http.Error(w, "Error while reading body", http.StatusInternalServerError)
		return
	}
	var comment database.SimpleComment
	decodingError := json.Unmarshal(bodyContent, &comment)
	if decodingError != nil {
		rt.baseLogger.WithError(decodingError).Error("Error while decoding body")
		http.Error(w, "Error while decoding body", http.StatusInternalServerError)
		return
	}
	commentError := rt.db.CommentPost(comment)
	if commentError != nil {
		rt.baseLogger.WithError(commentError).Error("Error while commenting post")
		http.Error(w, "Error while commenting post", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	_, writeError := w.Write([]byte("Post commented!"))
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
	}
}
