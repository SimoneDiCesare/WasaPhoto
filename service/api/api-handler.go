package api

import (
	"errors"
	"net/http"
)

/** TODO:
 * Add here endpoint -> rt.router.X("{path}", rt.FUNCTION)
 * Create rt.FUNCTION for each endpoint -> create file FUNCTION.go and write function there
 *		func (rt *_router) UNCTION(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
 * for DB Action -> create under database/ directory.
 */

/**
 * TODO: Missing Endpoints:
 * /users/:uid GET (getUserProfile)
 * /users/:uid/feeds GET (getMyStream)
 * /users/:uid/bans GET (getBans)
 * /users/:uid/bans/:bid PUT (banUser)
 * /users/:uid/bans/:bid DELETE (unbanUser)
 * /users/:uid/posts GET (getUserPosts)
 * /users/:uid/posts/:pid GET (getUserPost)
 * /posts POST (uploadPhoto)
 * /posts/:pid DELETE (deletPhoto)
 * /posts/:pid/likes/:uid PUT (likePhoto)
 * /posts/:pid/likes/:uid DELETE (unlikePhoto)
 * /posts/:pid/comments GET (getPostComments)
 * /posts/:pid/comments POST (commentPhoto)
 * /posts/:pid/comments/:cid DELET (uncommentPhoto)
 * =======================================================
 * TODO: Bans check on:
 * searchUsers, getFollows, followUser, unfollowUser, getFollowers, banUser, getUserPosts, getUserPost, likePhoto, unlikePhoto, getPostComments, commentPhoto, uncommentPhoto
 */

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/login", rt.login)
	rt.router.GET("/users", rt.searchUsers)
	rt.router.PUT("/users/:uid", rt.changeUserName)
	rt.router.GET("/users/:uid/follows", rt.getFollows)
	rt.router.GET("/users/:uid/followers", rt.getFollowers)
	rt.router.PUT("/users/:uid/followers/:fid", rt.followUser)
	rt.router.DELETE("/users/:uid/followers/:fid", rt.unfollowUser)

	// ==== TODO: Remove or Check importance ====
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}

func (rt *_router) HandleTokenError(err error, w http.ResponseWriter) {
	if errors.Is(err, &AuthenticationError{}) {
		rt.baseLogger.WithError(err).Error("Unauthorized")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	} else if errors.Is(err, &AuthenticationError{}) {
		rt.baseLogger.WithError(err).Error("Forbidden")
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}
