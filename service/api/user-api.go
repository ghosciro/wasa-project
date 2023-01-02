package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// get username from query
	username := r.URL.Query().Get("username")
	// get token from db
	token, err := rt.db.DoLogin(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func (rt *_router) getHome(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := r.Header.Get("token")
	username := r.URL.Query().Get("username")
	valid_username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if valid_username != username {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(" your are not the owner "))
		return
	}
	// get my stream from db
	stream, err := rt.db.GetMyStream(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stream)

}

func (rt *_router) getUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := r.URL.Query().Get("username")
	token := r.Header.Get("token")
	// get usernames from db
	usernames, err := rt.db.GetUsers(username)
	var users []string
	for _, user := range usernames {
		if rt.db.Isnotbanned(token, user) {
			users = append(users, user)
		}
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (rt *_router) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := ps.ByName("username")
	token := r.Header.Get("token")
	user_p, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if !rt.db.Isnotbanned(user_p, username) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are banned"))
		return
	}
	// get user from db
	user, err := rt.db.GetUserProfile(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}


	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (rt *_router) postUserOptions(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// read new username from query
	username := ps.ByName("username")
	new_username := r.URL.Query().Get("username")
	token := r.Header.Get("token")
	valid_username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if valid_username != username {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are not the owner"))
		return
	}
	// insert new username in db
	err = rt.db.SetMyUserName(username, new_username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(new_username)
}

func (rt *_router) postUserFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := ps.ByName("username")
	otherusername := ps.ByName("otherusername")
	token := r.Header.Get("token")
	valid_username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	if valid_username != username {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("You are not the owner"))
		return
	}
	if !rt.db.Isnotbanned(valid_username, otherusername) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("You are banned"))
		return
	}

	// post user following new username

	err = rt.db.FollowUser(username, otherusername)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	user, err := rt.db.GetUserProfile(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (rt *_router) deleteUserFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	username := ps.ByName("username")
	otherusername := ps.ByName("otherusername")
	token := r.Header.Get("token")
	valid_username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	if valid_username != username {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("You are not the owner"))
		return
	}
	if !rt.db.Isnotbanned(valid_username, otherusername) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("You are banned"))
		return
	}

	// delete user following new username

	err = rt.db.UnfollowUser(username, otherusername)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	user, err := rt.db.GetUserProfile(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (rt *_router) postUserBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := ps.ByName("username")
	otherusername := ps.ByName("otherusername")
	token := r.Header.Get("token")
	valid_username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	if valid_username != username {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("You are not the owner"))
		return
	}

	err = rt.db.BanUser(username, otherusername)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("User banned"))
}
func (rt *_router) getUserBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := ps.ByName("username")
	token := r.Header.Get("token")
	valid_username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if valid_username != username {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are not the owner"))
		return
	}
	// get banned users from db
	bannedusers, err := rt.db.GetBanned(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bannedusers)
}
func (rt *_router) deleteUserBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := ps.ByName("username")
	otherusername := ps.ByName("otherusername")
	token := r.Header.Get("token")
	valid_username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if valid_username != username {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are not the owner"))
		return
	}
	err = rt.db.UnbanUser(username, otherusername)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Unbanned"))
}
