package api

import (
	"encoding/json"
	"net/http"

	"io"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// read from the request the token and the photo id

	user := ps.ByName("username")

	token := r.Header.Get("token")
	valid_username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if !rt.db.Isnotbanned(valid_username, user) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are banned"))
		return
	}

	Picture, err := rt.db.GetUserPhotos(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// send the photo
	err = json.NewEncoder(w).Encode(Picture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (rt *_router) UploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	print("upload photo:")
	token := r.Header.Get("token")
	// read from the request the token and the photo id
	user := ps.ByName("username")
	tokenizeduser, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if tokenizeduser != user {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are not the owner of this account"))
		return
	}
	photo, err := io.ReadAll(r.Body)
	photo_r := string(photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// add the photo to the database
	id, err := rt.db.UploadPhoto(user, photo_r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// send the photo id
	json.NewEncoder(w).Encode(id)
	json.NewEncoder(w).Encode(user)
}

func (rt *_router) DeleteUserPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// read from the request the token and the photo id
	photo := ps.ByName("photoid")
	user := ps.ByName("username")
	token := r.Header.Get("token")
	tokenizeduser, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if tokenizeduser != user {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are not the owner of this account"))
		return
	}
	// delete the photo from the database
	err = rt.db.DeletePhoto(photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
func (rt *_router) GetUserPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// read from the request the token and the photo id
	user := ps.ByName("username")
	token := r.Header.Get("token")
	valid_username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if !rt.db.Isnotbanned(valid_username, user) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are banned"))
		return
	}
	photo := ps.ByName("photoid")
	// get the photo from the database
	Picture, err := rt.db.GetPhoto(photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// send the photo
	err = json.NewEncoder(w).Encode(Picture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
