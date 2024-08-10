package api

import (
	"errors"
	"net/http"

	schema "github.com/SimoneDiCesare/WasaPhoto/service/api/schemas"
)

/** TODO:
 * Add here endpoint -> rt.router.X("{path}", rt.FUNCTION)
 * Create rt.FUNCTION for each endpoint -> create file FUNCTION.go and write function there
 *		func (rt *_router) UNCTION(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
 * for DB Action -> create under database/ directory.
 */

/**
 * TODO: Missing Endpoints:
 * /posts/:pid DELETE (deletPhoto)
 * /posts/:pid/likes/:uid PUT (likePhoto)
 * /posts/:pid/likes/:uid DELETE (unlikePhoto)
 * /posts/:pid/comments GET (getPostComments)
 * /posts/:pid/comments POST (commentPhoto)
 * /posts/:pid/comments/:cid DELET (uncommentPhoto)
 * =======================================================
 * TODO: Bans check on:
 * getUserPosts, getUserPost, likePhoto, unlikePhoto, getPostComments, commentPhoto, uncommentPhoto
 */

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	// User Generic operations
	rt.router.POST("/login", rt.login)
	rt.router.GET("/users", rt.searchUsers)
	rt.router.PUT("/users/:uid", rt.changeUserName)
	rt.router.GET("/users/:uid", rt.getUserProfile)
	rt.router.GET("/users/:uid/feeds", rt.getMyStream)
	// Follow operations
	rt.router.GET("/users/:uid/follows", rt.getFollows)
	rt.router.GET("/users/:uid/followers", rt.getFollowers)
	rt.router.PUT("/users/:uid/followers/:fid", rt.followUser)
	rt.router.DELETE("/users/:uid/followers/:fid", rt.unfollowUser)
	// Ban operations
	rt.router.GET("/users/:uid/bans", rt.getBans)
	rt.router.PUT("/users/:uid/bans/:bid", rt.banUser)
	rt.router.DELETE("/users/:uid/bans/:bid", rt.unbanUser)
	// Post operations
	rt.router.PUT("/posts", rt.uploadPhoto)
	rt.router.GET("/users/:uid/posts", rt.getUserPosts)
	rt.router.DELETE("/posts/:pid", rt.deletePhoto)
	rt.router.GET("/users/:uid/posts/:pid", rt.getUserPost)

	// ==== TODO: Remove or Check importance ====
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

func (rt *_router) HandleTokenError(err error, w http.ResponseWriter) {
	if errors.Is(err, schema.ErrNoAuthentication) {
		rt.baseLogger.WithError(err).Error("Unauthorized")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	} else if errors.Is(err, schema.ErrNotAuthorized) {
		rt.baseLogger.WithError(err).Error("Forbidden")
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

func (rt *_router) HandleBanError(err error, w http.ResponseWriter) {
	if errors.Is(err, schema.ErrNotAuthorized) {
		rt.baseLogger.WithError(err).Error("Unauthorized")
		http.Error(w, "Unauthorized", http.StatusForbidden)
	} else {
		rt.baseLogger.WithError(err).Error("Error checking request validity")
		http.Error(w, "Error checking request validity", http.StatusInternalServerError)
	}
}
