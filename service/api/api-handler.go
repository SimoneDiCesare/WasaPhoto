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
	// User Follows and Followers
	rt.router.GET("/users/:uid/follows", rt.authTokenWrap(rt.getFollows))
	rt.router.POST("/users/:uid/follows/:uid2", rt.authUidWrap(rt.followUser))
	rt.router.DELETE("/users/:uid/follows/:uid2", rt.authUidWrap(rt.unfollowUser))
	rt.router.GET("/users/:uid/follower", rt.authTokenWrap(rt.getFollowers))
	// User Privacy (Bans)
	rt.router.GET("/users/:uid/bans", rt.authUidWrap(rt.getBans))
	rt.router.POST("/users/:uid/bans/:uid2", rt.authUidWrap(rt.banUser))
	rt.router.DELETE("/users/:uid/bans/:uid2", rt.authUidWrap(rt.unbanUser))
	// Post
	rt.router.POST("/posts", rt.authTokenWrap(rt.createPost))
	rt.router.GET("/posts/:pid", rt.authTokenWrap(rt.getPost))
	rt.router.DELETE("/posts/:pid", rt.authTokenWrap(rt.deletePost))
	rt.router.GET("/posts/:pid/image", rt.authTokenWrap(rt.getPostImage))
	// Post Comments
	rt.router.GET("/posts/:pid/comments", rt.authTokenWrap(rt.getPostComments))
	rt.router.POST("/posts/:pid/comments", rt.authTokenWrap(rt.commentPost))
	rt.router.GET("/posts/:pid/comments/:cid", rt.authTokenWrap(rt.getPostComment))
	rt.router.DELETE("/posts/:pid/comments/:cid", rt.authTokenWrap(rt.deletePostComment))
	// Post Likes
	rt.router.GET("/posts/:pid/likes", rt.authUidWrap(rt.getPostLikes))
	rt.router.POST("/posts/:pid/likes", rt.authUidWrap(rt.likePost))
	rt.router.DELETE("/posts/:pid/likes/:uid", rt.authUidWrap(rt.unlikePost))
	// TODO: rt.router.GET("/posts/:pid/owner", rt.authUidWrap(rt.getPostOwner))
	return rt.router
}
