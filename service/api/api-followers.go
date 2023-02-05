package api

import (
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve username from path

	uname := ps.ByName("user_name")

	// Get the list of followers
	followers, err := rt.db.GetUserFollowers(uname)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(followers))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error getting user followers")
		return
	}

	_, err = w.Write([]byte(followers))

	if err != nil {
		ctx.Logger.WithError(err).Error("error writing response")
	}

}

func (rt *_router) getUserFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve username from the path

	uname := ps.ByName("user_name")

	following, err := rt.db.GetUserFollowing(uname)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(following))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error getting user following")
		return
	}

	_, err = w.Write([]byte(following))
	if err != nil {
		ctx.Logger.WithError(err).Error("error writing response")
	}

}

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve target, username from path

	username := ps.ByName("user_name")
	followed_name := ps.ByName("followed_name")
	token := r.Header.Get("Authorization")

	is_valid, err := rt.db.Validate(username, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("error validating user"))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("invalid token"))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("invalid token")
		return
	}

	// Insert the follow relationship into the database

	ret_data, err := rt.db.FollowUser(username, followed_name)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		_, err := w.Write([]byte(ret_data))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error following user")

		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve username and target from path

	username := ps.ByName("user_name")
	followed_name := ps.ByName("followed_name")
	token := r.Header.Get("Authorization")

	is_valid, err := rt.db.Validate(username, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("error validating user"))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {

		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte("invalid token"))

		if err != nil {

			ctx.Logger.WithError(err).Error("error writing response")

		}

		ctx.Logger.WithError(err).Error("invalid token")
		return

	}

	// Insert the follow relationship into the database

	ret_data, err := rt.db.UnfollowUser(username, followed_name)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		_, err := w.Write([]byte(ret_data))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error unfollowing user")

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
