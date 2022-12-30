package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	rt.router.POST("/session", rt.wrap(rt.postSession))
	rt.router.GET("/home", rt.wrap(rt.getHome))
	rt.router.GET("/users", rt.wrap(rt.getUsers))
	rt.router.GET("/users/:username", rt.wrap(rt.getUser))
	rt.router.POST("/users/:username/options", rt.wrap(rt.postUserOptions))
	rt.router.POST("/users/:username/following/:otherusername", rt.wrap(rt.postUserFollowing))
	rt.router.DELETE("/users/:username/following/:otherusername", rt.wrap(rt.deleteUserFollowing))
	rt.router.POST("/users/:username/banned/:otherusername", rt.wrap(rt.postUserBanned))
	rt.router.DELETE("/users/:username/banned/:otherusername", rt.wrap(rt.deleteUserBanned))
	rt.router.GET("/users/:username/banned", rt.wrap(rt.getUserBanned))
	rt.router.POST("/users/:username/photos", rt.wrap(rt.UploadPhoto))
	rt.router.GET("/users/:username/photos", rt.wrap(rt.getUserPhotos))
	rt.router.DELETE("/users/:username/photos/:photoid", rt.wrap(rt.DeleteUserPhoto))
	rt.router.GET("/users/:username/photos/:photoid", rt.wrap(rt.GetUserPhoto))
	rt.router.PUT("/users/:username/photos/:photoid/likes", rt.wrap(rt.putUserPhotosLikes))
	rt.router.DELETE("/users/:username/photos/:photoid/likes", rt.wrap(rt.deleteUserPhotosLikes))
	rt.router.POST("/users/:username/photos/:photoid/comments", rt.wrap(rt.postUserPhotosComments))
	rt.router.DELETE("/users/:username/photos/:photoid/comments/:commentid", rt.wrap(rt.deleteUserPhotosComments))

	return rt.router
}
