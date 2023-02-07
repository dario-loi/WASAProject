package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/components"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// parse the request body to JSON
	decoder := json.NewDecoder(r.Body)
	var uname components.User
	err := decoder.Decode(&uname)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(components.BadRequestError))

		if err != nil {
			ctx.Logger.WithError(err).Error(
				fmt.Errorf("error writing response, details: %s", err).Error())
		}

		ctx.Logger.WithError(err).Error(
			fmt.Errorf("error parsing request body, details: %s", err).Error())

		return
	}

	// get the user ID
	ret_data, err := rt.db.PostUserID(uname.Uname)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(ret_data))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error(
			fmt.Errorf("error getting user ID (username: %s), details: %s", uname.Uname, err).Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(ret_data))

	if err != nil {
		ctx.Logger.WithError(err).Error(
			fmt.Errorf("error writing response, details: %s", err).Error())
	}

}
