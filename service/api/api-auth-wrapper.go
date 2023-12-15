package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// This wrap an handler for authenticated operation only.
// This are the ones that need only a valid token.
func (rt *_router) authTokenWrap(fn httprouter.Handle) func(http.ResponseWriter, *http.Request, httprouter.Params) {
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

// This wrap an handler for specific authenticated operation only.
// THis are the ones that need a valid token associated with the uid of the request
func (rt *_router) authUidWrap(fn httprouter.Handle) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Token")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		authError := rt.db.VerifyUidToken(ps.ByName("uid"), token)
		if authError != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fn(w, r, ps)
	}
}
