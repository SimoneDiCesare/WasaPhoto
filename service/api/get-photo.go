package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	dirPath := filepath.Join("uploads", pid)
	imagePath := filepath.Join(dirPath, "image.png")
	_, err := os.Stat(imagePath)
	if os.IsNotExist(err) {
		rt.baseLogger.WithError(err).Error("Photo does not exists")
		http.Error(w, "Photo does not exists", http.StatusNotFound)
		return
	}
	file, err := os.Open(imagePath)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Can't open file")
		http.Error(w, "Can't open file", http.StatusInternalServerError)
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			rt.baseLogger.WithError(err).Error("Can't close file")
		}
	}()
	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, imagePath)
}
