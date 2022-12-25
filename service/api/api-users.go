package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

type DummyUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

// GetUserProfile returns the profile of a user.
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get user id
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		data, err := json.MarshalIndent(Error{201, "Invalid user id"}, "", "  ")

		if err != nil {
			fmt.Print(err)
			w.Write([]byte("Internal server error, failed to marshal error message."))
			return
		}

		w.Write(data)
		return
	}

	user := &DummyUser{}

	users_res, err := rt.db.QueryAndPack(user, "SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := fmt.Sprintf("Internal server error, failed to query database: %s", err.Error())

		data, err := json.MarshalIndent(Error{202, msg}, "", "  ")

		if err != nil {
			fmt.Print(err)
			w.Write([]byte(
				fmt.Sprintf("Internal server error, failed to marshal error message, original error: %s",
					err.Error())))
			return
		}

		w.Write(data)
		return
	}

	if len(users_res) == 0 {
		w.WriteHeader(http.StatusNotFound)

		data, err := json.MarshalIndent(Error{203, "User not found"}, "", "  ")

		if err != nil {
			fmt.Print(err)
			w.Write([]byte("Internal server error, failed to marshal error message."))
			return
		}

		w.Write(data)
		return
	}

	user = users_res[0].(*DummyUser)

	data, err := json.MarshalIndent(user, "", "  ")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		data, err := json.MarshalIndent(Error{204, "Internal server error, failed to marshal user."}, "", "  ")

		if err != nil {
			fmt.Print(err)
			w.Write([]byte("Internal server error, failed to marshal error message."))
			return
		}

		w.Write(data)
		return
	}

	w.Write(data)

}
