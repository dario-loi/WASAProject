package api

import (
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
		fmt.Println("Empty request header")
		return
	}

	// Check if the user exists

	exists, err := rt.db.CheckUserExists(id)

	if err != nil || !exists {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Errorf(components.UnauthorizedErrorF, err).Error()))
		fmt.Println(fmt.Errorf("error authenticating: %w", err))
		return
	}
	// get the token from the request query

	json_out := r.URL.Query().Get("search_term")

	fmt.Printf("searching for %s, by %s\n", json_out, id)

	// get the list of users with the given name
	ret_data, err := rt.db.SearchUserByName(json_out)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ret_data))
		fmt.Println(fmt.Errorf("error getting user ID: %w", err))
		return
	}

	fmt.Println(ret_data)
	w.Write([]byte(ret_data))

}

func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get user id from path

	name := ps.ByName("user_name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(components.BadRequestError))
		fmt.Println("Empty user id")
		return
	}

	// Check if the user exists

	exists, err := rt.db.CheckUsernameExists(name)

	if err != nil || !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Errorf(components.NotFoundErrorF, err).Error()))
		fmt.Println("Requested user does not exist")
		return
	}

	// Get the user id

	id, err := rt.db.GetUserID(name)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(id))
		fmt.Println(fmt.Errorf("error getting user ID: %w", err))
		return
	}

	// Get the list of photos from the database

	ret_data, err := rt.db.GetUserPhotos(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ret_data))
		fmt.Println(fmt.Errorf("error getting user photos: %w", err))
		return
	}

	fmt.Println(ret_data)
	w.Write([]byte(ret_data))

}
