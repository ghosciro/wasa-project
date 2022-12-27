package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	rt.router.POST("/session", rt.wrap(rt.postSession))
	rt.router.GET("/home", rt.wrap(rt.getHome))
	rt.router.GET("/users", rt.wrap(rt.getUsers))
	rt.router.GET("/users/:id", rt.wrap(rt.getUser))
	rt.router.POST("/users/:id/options", rt.wrap(rt.postUserOptions))
	rt.router.POST("/users/:id/following/:otherusername", rt.wrap(rt.postUserFollowing))
	rt.router.DELETE("/users/:id/following/:otherusername", rt.wrap(rt.deleteUserFollowing))
	rt.router.POST("/users/:id/banned/:otherusername", rt.wrap(rt.postUserBanned))
	rt.router.DELETE("/users/:id/banned/:otherusername", rt.wrap(rt.deleteUserBanned))
	rt.router.GET("/users/:id/banned", rt.wrap(rt.getUserBanned))
	//rt.router.POST("users:id/Photos", rt.wrap(rt.postUserPhotos))
	//rt.router.GET("/users:id/Photos:id", rt.wrap(rt.getUserPhotos))
	//rt.router.DELETE("/users:id/Photos:id", rt.wrap(rt.deleteUserPhotos))
	//rt.router.PUT("/users:username/Photos:photoid/likes", rt.wrap(rt.putUserPhotosLikes))
	//rt.router.DELETE("/users:username/Photos:photoid/likes", rt.wrap(rt.deleteUserPhotosLikes))
	//rt.router.POST("/users:id/Photos:id/comments", rt.wrap(rt.postUserPhotosComments))
	//rt.router.DELETE("/users:id/Photos:id/comments:id", rt.wrap(rt.deleteUserPhotosComments))

	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
