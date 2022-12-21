package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	//get username from query
	var username string
	err := json.NewDecoder(r.Body).Decode(&username)
	//get token from db
	print(username + "\n")
	token, err := rt.db.DoLogin(username)
	print("token:", token)
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
	print(ps)
	username := ps.ByName("username")
	//get usernames from db
	print(username + "\n")
	usernames, err := rt.db.GetUsers("", username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	print("usernames:", usernames)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usernames)
}

func (rt *_router) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

}

func (rt *_router) postUserOptions(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
}

func (rt *_router) postUserFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
}

func (rt *_router) deleteUserFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
}

func (rt *_router) postUserBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
}

func (rt *_router) deleteUserBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
}
