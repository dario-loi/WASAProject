package api

import (
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserBans(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the user ID
	ret_data, err := rt.db.GetUserBans(ps.ByName("user_name"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ret_data))
		ctx.Logger.WithError(err).Error("error getting user bans")
		return
	}

	w.Write([]byte(ret_data))

}
