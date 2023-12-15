package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	// Special routes
	rt.router.GET("/liveness", rt.liveness)
	// Session
	rt.router.POST("/session", rt.postUserSession)
	// User
	rt.router.GET("/users", rt.authTokenWrap(rt.SearchUsers))
	rt.router.GET("/users/:uid", rt.authTokenWrap(rt.redirectoToUserProfile))
	rt.router.PUT("/users/:uid", rt.authUidWrap(rt.changeUserName))
	rt.router.DELETE("/users/:uid", rt.authUidWrap(rt.deleteUser))
	rt.router.GET("/users/:uid/profile", rt.authTokenWrap(rt.getUserProfile))
	rt.router.GET("/users/:uid/image", rt.authTokenWrap(rt.getUserProfileImage))
	rt.router.PUT("/users/:uid/image", rt.authUidWrap(rt.changeUserProfileImage))
	// Follows
	rt.router.GET("/users/:uid/follows", rt.authTokenWrap(rt.getFollows))
	rt.router.POST("/users/:uid/follows/:uid2", rt.authUidWrap(rt.followUser))
	rt.router.DELETE("/users/:uid/follows/:uid2", rt.authUidWrap(rt.unfollowUser))
	rt.router.GET("/users/:uid/follower", rt.authTokenWrap(rt.getFollowers))
	// Privacy
	rt.router.GET("/users/:uid/bans", rt.authUidWrap(rt.getBans))
	rt.router.POST("/users/:uid/bans/:uid2", rt.authUidWrap(rt.banUser))
	rt.router.DELETE("/users/:uid/bans/:uid2", rt.authUidWrap(rt.unbanUser))
	// Post
	rt.router.POST("/posts", rt.authTokenWrap(rt.createPost))
	rt.router.GET("/posts/:pid", rt.authTokenWrap(rt.getPost))
	rt.router.DELETE("/posts/:pid", rt.authTokenWrap(rt.deletePost))
	rt.router.GET("/posts/:pid/image", rt.authTokenWrap(rt.getPostImage))
	rt.router.GET("/posts/:pid/comments", rt.authTokenWrap(rt.getPostComments))
	rt.router.POST("/posts/:pid/comments", rt.authTokenWrap(rt.commentPost))
	//TODO:
	//		GET comment of uid on pid
	//		DELETE comment of uid on pid
	//		GET post's uid
	//		POST add like to pid from uid
	//		GET likes of pid
	//		DELETE like of uid on pid
	rt.router.GET("/posts/:pid/comments/:cid", rt.authTokenWrap(rt.getPostComment))
	rt.router.DELETE("/posts/:pid/comments/:cid", rt.authUidWrap(rt.deletePostComment))
	rt.router.GET("/posts/:pid/owner", rt.authUidWrap(rt.getPostOwner))
	rt.router.GET("/posts/:pid/likes", rt.authUidWrap(rt.getPostLikes))
	rt.router.POST("/posts/:pid/likes", rt.authUidWrap(rt.likePost))
	rt.router.DELETE("/posts/:pid/likes/:uid", rt.authUidWrap(rt.unlikePost))
	return rt.router
}
