package api

import (
	"io"
	"net/http"

	"github.com/SimoneDiCesare/WasaPhoto/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfileImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	profileImagePath := database.GetImage("./users/" + uid + "/image.png")
	imageFile, openError := http.Dir("").Open(profileImagePath)
	if openError != nil {
		rt.baseLogger.WithError(openError).Error("Error opening user profile image")
		http.Error(w, "Error opening user profile image", http.StatusInternalServerError)
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
		rt.baseLogger.WithError(writeError).Error("Error writing user profile image")
		http.Error(w, "Error writing user profile image", http.StatusInternalServerError)
	}
}
