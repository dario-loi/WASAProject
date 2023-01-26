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

	// User routes

	rt.router.GET("/users", rt.wrap(rt.searchUser))

	// Session routes
	rt.router.PUT("/session", rt.wrap(rt.doLogin))

	// Getters
	rt.router.GET("/resources/photos/:UUID", rt.wrap(rt.getPhoto))

	rt.router.GET("/users/:user_name/profile/photos", rt.wrap(rt.getUserPhotos))
	rt.router.GET("/users/:user_name/profile", rt.wrap(rt.getUserProfile))

	rt.router.GET("/users/:user_name/followers", rt.wrap(rt.getUserFollowers))
	rt.router.GET("/users/:user_name/following", rt.wrap(rt.getUserFollowing))

	rt.router.GET("/users/:user_name/profile/photos/:photo_id/likes",
		rt.wrap(rt.GetPhotoLikes))

	rt.router.GET("/users/:user_name/profile/photos/:photo_id/comments",
		rt.wrap(rt.GetPhotoComments))

	rt.router.GET("/users/:user_name/bans", rt.wrap(rt.getUserBans))

	// Follower routes

	rt.router.PUT("/users/:user_name/following/:followed_name", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:user_name/following/:followed_name", rt.wrap(rt.unfollowUser))

	// Ban routes

	rt.router.PUT("/users/:user_name/bans/:banned_name", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:user_name/bans/:banned_name", rt.wrap(rt.unbanUser))

	// Like routes

	rt.router.PUT("/users/:user_name/profile/photos/:photo_id/likes/:liker_id", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:user_name/profile/photos/:photo_id/likes/:liker_id", rt.wrap(rt.unlikePhoto))

	// Comment routes

	rt.router.PUT("/users/:user_name/profile/photos/:photo_id/comments/:comment_id", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:user_name/profile/photos/:photo_id/comments/:comment_id", rt.wrap(rt.deleteComment))

	// Photo routes

	rt.router.PUT("/users/:user_name/profile/photos", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:user_name/profile/photos/:photo_id", rt.wrap(rt.deletePhoto))

	// Username change routes

	rt.router.PUT("/users/:user_name/username", rt.wrap(rt.changeUsername))

	return rt.router
}
