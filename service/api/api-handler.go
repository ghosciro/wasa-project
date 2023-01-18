package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	rt.router.GET("/", rt.wrap(rt.isalive))
	rt.router.GET("/home", rt.wrap(rt.getHome))
	rt.router.GET("/users", rt.wrap(rt.getUsers))
	rt.router.GET("/users/:username", rt.wrap(rt.getUser))
	rt.router.GET("/users/:username/banned", rt.wrap(rt.getUserBanned))
	rt.router.GET("/users/:username/Photos", rt.wrap(rt.getUserPhotos))
	rt.router.GET("/users/:username/Photos/:photoid", rt.wrap(rt.GetUserPhoto))
	rt.router.GET("/users/:username/Photos/:photoid/likes", rt.wrap(rt.getUserPhotosLikes))
	rt.router.GET("/users/:username/Photos/:photoid/comments", rt.wrap(rt.getUserPhotosComments))

	rt.router.POST("/users/:username/following/:otherusername", rt.wrap(rt.postUserFollowing))
	rt.router.POST("/session", rt.wrap(rt.postSession))
	rt.router.POST("/users/:username/Photos", rt.wrap(rt.UploadPhoto))
	rt.router.POST("/users/:username/banned/:otherusername", rt.wrap(rt.postUserBanned))
	rt.router.POST("/users/:username/Photos/:photoid/comments", rt.wrap(rt.postUserPhotosComments))

	rt.router.PUT("/users/:username/options", rt.wrap(rt.postUserOptions))
	rt.router.PUT("/users/:username/Photos/:photoid/likes", rt.wrap(rt.putUserPhotosLikes))

	rt.router.DELETE("/session", rt.wrap(rt.deleteSession))
	rt.router.DELETE("/users/:username/following/:otherusername", rt.wrap(rt.deleteUserFollowing))
	rt.router.DELETE("/users/:username/banned/:otherusername", rt.wrap(rt.deleteUserBanned))
	rt.router.DELETE("/users/:username/Photos/:photoid", rt.wrap(rt.DeleteUserPhoto))
	rt.router.DELETE("/users/:username/Photos/:photoid/likes/:likeid", rt.wrap(rt.deleteUserPhotosLikes))
	rt.router.DELETE("/users/:username/Photos/:photoid/comments/:commentid", rt.wrap(rt.deleteUserPhotosComments))

	return rt.router
}
