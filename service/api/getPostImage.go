package api

import (
	"io"
	"net/http"

	"github.com/SimoneDiCesare/WasaPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPostImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	postImagePath := database.GetImage("/posts/" + pid + "/image.png")
	imageFile, openError := http.Dir(".").Open(postImagePath)
	if openError != nil {
		rt.baseLogger.WithError(openError).Error("Error opening post image")
		http.Error(w, "Error opening post image", http.StatusInternalServerError)
		return
	}
	defer func() {
		closeError := imageFile.Close()
		if closeError != nil {
			rt.baseLogger.WithError(closeError).Error("Error closing image stream")
		}
	}()
	w.Header().Set("Content-Type", "image/png")
	_, writeError := io.Copy(w, imageFile)
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error writing post image")
		http.Error(w, "Error writing post image", http.StatusInternalServerError)
	}
}
