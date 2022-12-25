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
	var uname components.Username
	err := decoder.Decode(&uname)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(components.BadRequestError))
		fmt.Println(fmt.Errorf("error decoding JSON: %w", err))
		return
	}

	// get the user ID
	ret_data, err := rt.db.PostUserID(uname.Username_string)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ret_data))
		fmt.Println(fmt.Errorf("error getting user ID: %w", err))
		return
	}

	w.Write([]byte(ret_data))

}
