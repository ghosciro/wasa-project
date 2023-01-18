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
	token := r.Header.Get("Authorization")
	otheruser := ps.ByName("username")
	valid_user, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !rt.db.Isnotbanned(valid_user, otheruser) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := rt.db.GetUserToken(token)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	photo := ps.ByName("photoid")
	err = rt.db.LikePhoto(user, photo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (rt *_router) deleteUserPhotosLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := r.Header.Get("Authorization")
	user, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	photo := ps.ByName("photoid")
	err = rt.db.UnlikePhoto(user, photo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (rt *_router) getUserPhotosComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photo := ps.ByName("photoid")
	token := r.Header.Get("Authorization")
	otheruser := ps.ByName("username")
	valid_user, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !rt.db.Isnotbanned(valid_user, otheruser) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	comments, err := rt.db.GetComments(photo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		return
	}
}

func (rt *_router) getUserPhotosLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photo := ps.ByName("photoid")
	token := r.Header.Get("Authorization")
	otheruser := ps.ByName("username")
	valid_user, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !rt.db.Isnotbanned(valid_user, otheruser) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	likes, err := rt.db.GetLikes(photo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(likes)
	if err != nil {
		return
	}
}
func (rt *_router) postUserPhotosComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photo := ps.ByName("photoid")
	token := r.Header.Get("Authorization")
	otheruser := ps.ByName("username")
	valid_user, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !rt.db.Isnotbanned(valid_user, otheruser) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	comment, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	comment_s := string(comment)
	id, err := rt.db.CommentPhoto(valid_user, photo, comment_s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(id)
	if err != nil {
		return
	}

}
func (rt *_router) deleteUserPhotosComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photo := ps.ByName("photoid")
	token := r.Header.Get("Authorization")
	user, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	comment_id, err := strconv.Atoi(ps.ByName("commentid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.UncommentPhoto(user, photo, comment_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}
