package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) authWrap(fn httprouter.Handle) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Token")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		authError := rt.db.VerifyToken(token)
		if authError != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fn(w, r, ps)
	}
}
