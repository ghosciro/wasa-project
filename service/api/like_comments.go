package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) putUserPhotosLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := r.Header.Get("token")
	otheruser := ps.ByName("username")
	valid_user, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	if !rt.db.Isnotbanned(valid_user, otheruser) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are banned"))
		return
	}
	user, err := rt.db.GetUserToken(token)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	photo := ps.ByName("photoid")
	err = rt.db.LikePhoto(user, photo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (rt *_router) deleteUserPhotosLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := r.Header.Get("token")
	user, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	photo := ps.ByName("photoid")
	err = rt.db.UnlikePhoto(user, photo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (rt *_router) postUserPhotosComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photo := ps.ByName("photoid")
	token := r.Header.Get("token")
	otheruser := ps.ByName("username")
	valid_user, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	if !rt.db.Isnotbanned(valid_user, otheruser) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("You are banned"))
		return
	}
	comment, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	comment_s := string(comment)
	id, err := rt.db.CommentPhoto(valid_user, photo, comment_s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(id)

}
func (rt *_router) deleteUserPhotosComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photo := ps.ByName("photoid")
	token := r.Header.Get("token")
	user, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ =w.Write([]byte(err.Error()))
		return
	}
	comment_id, err := strconv.Atoi(ps.ByName("commentid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	err = rt.db.UncommentPhoto(user, photo, comment_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)

}
