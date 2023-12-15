package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPostComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	_, writeError := w.Write([]byte("Post Comments retrieved!"))
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error while writing response", http.StatusInternalServerError)
	}
}
