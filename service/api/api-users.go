package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/components"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	id := r.Header.Get("searcher_id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(components.BadRequestError))
		ctx.Logger.Error("Empty request header")
		return
	}

	// Check if the user exists

	exists, err := rt.db.CheckUserExists(id)

	if err != nil || !exists {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Errorf(components.UnauthorizedErrorF, err).Error()))
		return
	}
	// get the token from the request query

	json_out := r.URL.Query().Get("search_term")

	// get the list of users with the given name
	ret_data, err := rt.db.SearchUserByName(json_out)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ret_data))
		ctx.Logger.WithError(err).Error("error searching user")
		return
	}

	w.Write([]byte(ret_data))

}

func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get username from path

	name := ps.ByName("user_name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(components.BadRequestError))
		ctx.Logger.Error("Empty username")
		return
	}

	// Check if the user exists

	exists, err := rt.db.CheckUsernameExists(name)

	if err != nil || !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Errorf(components.NotFoundErrorF, err).Error()))
		ctx.Logger.Error("error checking user existence")
		return
	}

	// Get the user id

	id, err := rt.db.GetUserID(name)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(id))
		ctx.Logger.WithError(err).Error("error getting user id")
		return
	}

	// Get the list of photos from the database

	ret_data, err := rt.db.GetUserPhotos(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ret_data))
		ctx.Logger.WithError(err).Error("error getting user photos")
		return
	}

	w.Write([]byte(ret_data))

}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get user name from path

	name := ps.ByName("user_name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(components.BadRequestError))
		ctx.Logger.Error("Empty username")
		return
	}

	// Check if the user exists

	exists, err := rt.db.CheckUsernameExists(name)

	if err != nil || !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Errorf(components.NotFoundErrorF, err).Error()))
		ctx.Logger.Error("error checking user existence")
		return
	}

	// Get the user id

	id, err := rt.db.GetUserID(name)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(id))
		ctx.Logger.WithError(err).Error("error getting user id")
		return
	}

	// Get the user profile from the database

	ret_data, err := rt.db.GetUserProfile(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ret_data))
		ctx.Logger.WithError(err).Error("error getting user profile")
		return
	}

	w.Write([]byte(ret_data))

}

func (rt *_router) changeUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get user name from path

	user_name := ps.ByName("user_name")

	// Get user id from header

	id := r.Header.Get("user_id")

	is_valid, err := rt.db.Validate(id, user_name)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Errorf(components.UnauthorizedErrorF, err).Error()))
		ctx.Logger.WithError(err).Error("error authenticating")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(components.UnauthorizedError))
		ctx.Logger.Error("error authenticating")
		return
	}

	// Get the new username from the request body

	// Read the request body

	dec := json.NewDecoder(r.Body)

	var new_username string

	err = dec.Decode(&new_username)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(components.BadRequestError))
		ctx.Logger.WithError(err).Error("error decoding request body")
		return
	}

	// Change the username in the database

	ret_data, err := rt.db.ChangeUsername(user_name, new_username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ret_data))
		ctx.Logger.WithError(err).Error("error changing username")
		return
	}

	w.Write([]byte(ret_data))

}
