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
		_, err := w.Write([]byte(ret_data))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error getting user bans")
		return
	}

	_, err = w.Write([]byte(ret_data))

	if err != nil {
		ctx.Logger.WithError(err).Error("error writing response")
	}

}

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the user ID
	token := r.Header.Get("user_id")
	banisher := ps.ByName("user_name")

	is_valid, err := rt.db.Validate(banisher, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	to_ban := ps.ByName("banned_name")

	ret, err := rt.db.BanUser(banisher, to_ban)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error banning user")
		_, err := w.Write([]byte(ret))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the user ID
	token := r.Header.Get("user_id")
	banisher := ps.ByName("user_name")

	is_valid, err := rt.db.Validate(banisher, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	to_unban := ps.ByName("banned_name")

	ret, err := rt.db.UnbanUser(banisher, to_unban)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error unbanning user")
		_, err := w.Write([]byte(ret))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)

}
