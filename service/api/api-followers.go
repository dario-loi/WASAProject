package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/components"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve username from request body

	decoder := json.NewDecoder(r.Body)
	var uname components.User
	err := decoder.Decode(&uname)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(components.BadRequestError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error decoding JSON")
		return
	}

	// Get the list of followers
	followers, err := rt.db.GetUserFollowers(uname.Uname)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(followers))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error getting user followers")
		return
	}

	// Unmarshal the followers JSON into a slice of names

	var followers_names []string
	err = json.Unmarshal([]byte(followers), &followers_names)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(followers))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error unmarshaling followers JSON")
		return
	}

	return_struct := struct {
		Username  string   `json:"owner"`
		Followers []string `json:"follow-list"`
	}{
		Username:  uname.Uname,
		Followers: followers_names,
	}

	// Marshal the struct into JSON

	ret_data, err := json.Marshal(return_struct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(followers))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error marshaling followers JSON")
		return
	}

	_, err = w.Write(ret_data)

	if err != nil {
		ctx.Logger.WithError(err).Error("error writing response")
	}

}

func (rt *_router) getUserFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve username from request body

	decoder := json.NewDecoder(r.Body)
	var uname components.User
	err := decoder.Decode(&uname)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		_, err := w.Write([]byte(components.BadRequestError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error decoding JSON")

		return
	}

	following, err := rt.db.GetUserFollowing(uname.Uname)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(following))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error getting user following")
		return
	}

	// Unmarshal the followers JSON into a slice of names

	var following_names []string
	err = json.Unmarshal([]byte(following), &following_names)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		_, err := w.Write([]byte(following))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error unmarshaling following JSON")

		return
	}

	return_struct := struct {
		Username  string   `json:"owner"`
		Following []string `json:"follow-list"`
	}{
		Username:  uname.Uname,
		Following: following_names,
	}

	// Marshal the struct into JSON

	ret_data, err := json.Marshal(return_struct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(following))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error marshaling following JSON")
		return
	}

	_, err = w.Write(ret_data)
	if err != nil {
		ctx.Logger.WithError(err).Error("error writing response")
	}

}

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve username and target from header

	username := r.Header.Get("user_name")
	followed_name := r.Header.Get("followed_name")
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

	// Retrieve username and target from header

	username := r.Header.Get("user_name")
	followed_name := r.Header.Get("followed_name")
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
