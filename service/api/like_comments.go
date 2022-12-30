package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) putUserPhotosLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	user := ps.ByName("username")
	photo := ps.ByName("photoid")
	print(user, photo)
	err := rt.db.LikePhoto(user, photo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (rt *_router) deleteUserPhotosLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	user := ps.ByName("username")
	photo := ps.ByName("photoid")
	err := rt.db.UnlikePhoto(user, photo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (rt *_router) postUserPhotosComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photo := ps.ByName("photoid")
	user := ps.ByName("username")
	comment, err := ioutil.ReadAll(r.Body)
	comment_s := string(comment)
	id, err := rt.db.CommentPhoto(user, photo, comment_s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}
func (rt *_router) deleteUserPhotosComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photo := ps.ByName("photoid")
	user := ps.ByName("username")
	comment_id, err := strconv.Atoi(ps.ByName("commentid"))
	print(photo, user, comment_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = rt.db.UncommentPhoto(user, photo, comment_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)

}
