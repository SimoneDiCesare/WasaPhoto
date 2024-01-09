package api

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) changeUserProfileImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	parseError := r.ParseMultipartForm(10 << 20)
	if parseError != nil {
		rt.baseLogger.WithError(parseError).Error("Error parsing image")
		http.Error(w, "Error parsing image", http.StatusInternalServerError)
		return
	}
	newImage, _, openFileError := r.FormFile("image")
	if openFileError != nil {
		rt.baseLogger.WithError(openFileError).Error("Error opening image file")
		http.Error(w, "Error opening image file", http.StatusInternalServerError)
		return
	}
	defer func() {
		closeError := newImage.Close()
		if closeError != nil {
			rt.baseLogger.WithError(closeError).Error("Error closing image file")
		}
	}()
	imageDir := "./.users/" + uid
	preparingError := os.MkdirAll(imageDir, os.ModePerm)
	if preparingError != nil {
		rt.baseLogger.WithError(preparingError).Error("Error preparing image dir")
		http.Error(w, "Error preparing image dir", http.StatusInternalServerError)
		return
	}
	imagePath := filepath.Join(imageDir, "image.png")
	newFile, createError := os.Create(imagePath)
	defer func() {
		closeError := newFile.Close()
		if closeError != nil {
			rt.baseLogger.WithError(closeError).Error("Error closing image file")
		}
	}()
	if createError != nil {
		rt.baseLogger.WithError(createError).Error("Error creating image")
		http.Error(w, "Error creating image", http.StatusInternalServerError)
		return
	}
	_, copyError := io.Copy(newFile, newImage)
	if copyError != nil {
		rt.baseLogger.WithError(copyError).Error("Error saving image")
		http.Error(w, "Error saving image", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	_, writeError := w.Write([]byte("Profile image changed!"))
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error while writing response")
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
