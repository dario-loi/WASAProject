package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/components"
	"github.com/julienschmidt/httprouter"
)

// func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

// 	// parse the request body to JSON
// 	decoder := json.NewDecoder(r.Body)
// 	var uname components.Username
// 	err := decoder.Decode(&uname)

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(components.BadRequestError))
// 		ctx.Logger().WithError(err).Error("error decoding JSON")
// 		return
// 	}

// 	// get the user ID
// 	ret_data, err := rt.db.PostUserID(uname.Username_string)

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(ret_data))
// 		ctx.Logger().WithError(err).Error("error getting user ID")
// 		return
// 	}

// 	w.Write([]byte(ret_data))

// }

func (rt *_router) getUserFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve username from request body

	decoder := json.NewDecoder(r.Body)

	var uname components.User

	err := decoder.Decode(&uname)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		w.Write([]byte(components.BadRequestError))

		ctx.Logger.WithError(err).Error("error decoding JSON")

		return // exit the function
	}

	// Get the list of followers

	followers, err := rt.db.GetUserFollowers(uname.Uname)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(followers))

		ctx.Logger.WithError(err).Error("error getting user followers")

		return
	}

	// Unmarshal the followers JSON into a slice of names

	var followers_names []string

	err = json.Unmarshal([]byte(followers), &followers_names)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(followers))

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

		w.Write([]byte(followers))

		ctx.Logger.WithError(err).Error("error marshaling followers JSON")

		return
	}

	w.Write(ret_data)

}

func (rt *_router) getUserFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve username from request body

	decoder := json.NewDecoder(r.Body)

	var uname components.User

	err := decoder.Decode(&uname)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		w.Write([]byte(components.BadRequestError))

		ctx.Logger.WithError(err).Error("error decoding JSON")

		return // exit the function
	}

	// Get the list of followers

	following, err := rt.db.GetUserFollowing(uname.Uname)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(following))

		ctx.Logger.WithError(err).Error("error getting user following")

		return
	}

	// Unmarshal the followers JSON into a slice of names

	var following_names []string

	err = json.Unmarshal([]byte(following), &following_names)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(following))

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

		w.Write([]byte(following))

		ctx.Logger.WithError(err).Error("error marshaling following JSON")

		return
	}

	w.Write(ret_data)

}
