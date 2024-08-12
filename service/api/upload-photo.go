package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse multipart body -> 20 MB
	uid, tokenErr := rt.checkToken(r.Header.Get("token"))
	if tokenErr != nil {
		rt.HandleTokenError(tokenErr, w)
		return
	}
	err := r.ParseMultipartForm(20 << 20)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Error parsing body")
		http.Error(w, "Error parsing body", http.StatusBadRequest)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		rt.baseLogger.WithError(err).Error("Can't read file")
		http.Error(w, "Can't read file", http.StatusBadRequest)
		return
	}
	defer func() { // Close Body File
		closeError := file.Close()
		if closeError != nil {
			rt.baseLogger.WithError(closeError).Error("Can't close file. May cause issue in future")
		}
	}()
	// Crea un file temporaneo per salvare il file caricato
	pid, err := rt.db.CreatePost(uid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Can't create post on db")
		http.Error(w, "Can't create post on db", http.StatusInternalServerError)
		return
	}
	photoFile, err := os.Create("uploads/" + uid + "/" + pid + ".png")
	if err != nil {
		rt.baseLogger.WithError(err).Error("Can't create file on storage")
		http.Error(w, "Can't create file on storage", http.StatusInternalServerError)
		return
	}
	defer func() { // Close Photo File
		closeError := photoFile.Close()
		if closeError != nil {
			rt.baseLogger.WithError(closeError).Error("Can't close file. May cause issue in future")
		}
	}()
	_, err = io.Copy(photoFile, file)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Can't save image on storage")
		http.Error(w, "Can't save image on storage", http.StatusInternalServerError)
		return
	}
	simplePost, err := rt.db.GetSimplePost(pid)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Can't retrieve post data")
		http.Error(w, "Can't retrieve post data", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(simplePost)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Can't encode response")
		http.Error(w, "Can't encode response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		rt.baseLogger.WithError(err).Error("Can't write response")
		http.Error(w, "Can't write response", http.StatusInternalServerError)
		return
	}
}
