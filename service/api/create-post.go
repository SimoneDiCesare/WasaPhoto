package api

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) createPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	parseError := r.ParseMultipartForm(10 << 20)
	if parseError != nil {
		rt.baseLogger.WithError(parseError).Error("Error parsing request body")
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}
	uid, caption := r.FormValue("uid"), r.FormValue("caption")
	post, createError := rt.db.CreatePost(uid, caption)
	if createError != nil {
		rt.baseLogger.WithError(createError).Error("Error creating post")
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}
	// Writing Image
	imageData, _, formError := r.FormFile("image")
	if formError != nil {
		rt.baseLogger.WithError(formError).Error("Error parsing image data")
		http.Error(w, "Error parsing image data", http.StatusInternalServerError)
		return
	}
	defer func() {
		closeError := imageData.Close()
		if closeError != nil {
			rt.baseLogger.WithError(closeError).Error("Error closing file")
		}
	}()
	imageDir := ".posts/" + post.Pid
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
	_, copyError := io.Copy(newFile, imageData)
	if copyError != nil {
		rt.baseLogger.WithError(copyError).Error("Error saving image")
		http.Error(w, "Error saving image", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	_, writeError := w.Write([]byte("Post created!"))
	if writeError != nil {
		rt.baseLogger.WithError(writeError).Error("Error writing response")
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
