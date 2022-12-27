package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//read from the request the token and the photo id
	user := ps.ByName("user")
	photo := ps.ByName("photoid")

	//check if the user is banned
	banned, err := rt.checkban(user, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if banned {
		http.Error(w, "user is banned", http.StatusBadRequest)
		return
	}
	//
	Picture, err := rt.db.GetPhoto(photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//send the photo
	err = json.NewEncoder(w).Encode(Picture)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (rt *_router) UploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//read from the request the token and the photo id
	user := ps.ByName("user")
	var photo string
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//add the photo to the database
	id, err := rt.db.UploadPhoto(user, photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//send the photo id
	err = json.NewEncoder(w).Encode(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (rt *_router) checkban(token string, otheruserid string) (bool, error) {
	banned_users, err := rt.db.GetBanned(otheruserid)
	if err != nil {
		return true, err
	}
	for _, user := range banned_users {
		if user == token {
			return true, nil
		}
	}
	return false, nil
}
