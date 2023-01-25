package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/components"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetPhotoLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve photo ID from request body

	decoder := json.NewDecoder(r.Body)

	var photoID components.SHA256hash

	err := decoder.Decode(&photoID)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		w.Write([]byte(components.BadRequestError))

		ctx.Logger.WithError(err).Error("error decoding JSON")

		return // exit the function
	}

	// get the photo likes
	ret_data, err := rt.db.GetPhotoLikes(photoID.Hash)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ret_data))
		ctx.Logger.WithError(err).Error("error getting photo likes")
		return
	}

	w.Write([]byte(ret_data))

}

func (rt *_router) GetPhotoComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve photo ID from request body

	decoder := json.NewDecoder(r.Body)

	var photoID components.SHA256hash

	err := decoder.Decode(&photoID)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		w.Write([]byte(components.BadRequestError))

		ctx.Logger.WithError(err).Error("error decoding JSON")

		return // exit the function
	}

	// get the photo comments
	ret_data, err := rt.db.GetPhotoComments(photoID.Hash)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ret_data))
		ctx.Logger.WithError(err).Error("error getting photo comments")
		return
	}

	w.Write([]byte(ret_data))

}
