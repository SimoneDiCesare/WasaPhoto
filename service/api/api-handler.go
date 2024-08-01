package api

import (
	"net/http"
)

/** TODO:
 * Add here endpoint -> rt.router.X("{path}", rt.FUNCTION)
 * Create rt.FUNCTION for each endpoint -> create file FUNCTION.go and write function there
 *		func (rt *_router) UNCTION(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
 * for DB Action -> create under database/ directory.
 */

// TODO: Add here endpoints

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/login", rt.login)

	// ==== TODO: Remove or Check importance ====
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
