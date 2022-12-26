package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//get username from query
	username := r.URL.Query().Get("username")
	//get token from db
	token, err := rt.db.DoLogin(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func (rt *_router) getHome(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

}

func (rt *_router) getUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := r.URL.Query().Get("username")
	//get usernames from db
	usernames, err := rt.db.GetUsers("", username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usernames)
}

func (rt *_router) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := ps.ByName("id")
	print(username + "\n")
	//get user from db
	user, err := rt.db.GetUserProfile("", username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//print user

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (rt *_router) postUserOptions(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//read new username from query
	new_username := r.URL.Query().Get("username")
	//insert new username in db
	token, err := rt.db.SetMyUserName("", new_username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func (rt *_router) postUserFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
}

func (rt *_router) deleteUserFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
}

func (rt *_router) postUserBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
}

func (rt *_router) deleteUserBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
}
