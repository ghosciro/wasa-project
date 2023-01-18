package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) isalive(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.WriteHeader(http.StatusOK)
}

func (rt *_router) postSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// get username from body
	var username struct {
		Username string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&username)
	print("username:", username.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if username.Username == "null" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// get token from db
	token, err := rt.db.DoLogin(username.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (rt *_router) getHome(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := r.Header.Get("token")
	username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// get my stream from db
	stream, err := rt.db.GetMyStream(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(stream)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (rt *_router) deleteSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := r.Header.Get("token")
	err := rt.db.DoLogout(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func (rt *_router) getUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := r.URL.Query().Get("username")
	token := r.Header.Get("token")
	user_p, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// get usernames from db
	usernames, err := rt.db.GetUsers(username)

	var users []string
	for _, user := range usernames {
		if rt.db.Isnotbanned(user_p, user) {
			users = append(users, user)
		}
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (rt *_router) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := ps.ByName("username")
	token := r.Header.Get("token")
	user_p, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !rt.db.Isnotbanned(user_p, username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// get user from db
	user, err := rt.db.GetUserProfile(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (rt *_router) postUserOptions(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// read new username from query
	new_username := r.URL.Query().Get("username")
	token := r.Header.Get("token")
	username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if username != username {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// insert new username in db
	err = rt.db.SetMyUserName(username, new_username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(new_username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (rt *_router) postUserFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	otherusername := ps.ByName("otherusername")
	token := r.Header.Get("token")
	username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !rt.db.Isnotbanned(username, otherusername) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// post user following new username
	err = rt.db.FollowUser(username, otherusername)
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
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (rt *_router) deleteUserFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	otherusername := ps.ByName("otherusername")
	token := r.Header.Get("token")
	username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !rt.db.Isnotbanned(username, otherusername) {
		w.WriteHeader(http.StatusBadRequest)
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
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (rt *_router) postUserBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	otherusername := ps.ByName("otherusername")
	token := r.Header.Get("token")
	username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if username != username {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.BanUser(username, otherusername)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (rt *_router) getUserBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	token := r.Header.Get("token")
	username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if username != username {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// get banned users from db
	bannedusers, err := rt.db.GetBanned(username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(bannedusers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
func (rt *_router) deleteUserBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	otherusername := ps.ByName("otherusername")
	token := r.Header.Get("token")
	username, err := rt.db.GetUserToken(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if username != username {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.UnbanUser(username, otherusername)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, error := w.Write([]byte("Unbanned"))
	if error != nil {
		return
	}
}
