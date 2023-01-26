package api

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve photo ID from path

	photoID := ps.ByName("photoID")

	// get the user ID

	token := r.Header.Get("user_id")

	userName := ps.ByName("user_name")

	is_valid, err := rt.db.Validate(userName, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ret, err := rt.db.LikePhoto(userName, photoID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error liking photo")
		w.Write([]byte(ret))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve photo ID from path

	photoID := ps.ByName("photoID")

	// get the user ID

	token := r.Header.Get("user_id")

	userName := ps.ByName("user_name")

	is_valid, err := rt.db.Validate(userName, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ret, err := rt.db.UnlikePhoto(userName, photoID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error unliking photo")
		w.Write([]byte(ret))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve photo ID from path

	photoID := ps.ByName("photoID")

	// get the user ID

	token := r.Header.Get("user_id")

	userName := ps.ByName("user_name")

	is_valid, err := rt.db.Validate(userName, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Retrieve comment from request body

	decoder := json.NewDecoder(r.Body)

	var comment components.Comment

	err = decoder.Decode(&comment)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		w.Write([]byte(components.BadRequestError))

		ctx.Logger.WithError(err).Error("error decoding JSON")

		return // exit the function
	}

	comment_id := ps.ByName("comment_id")

	comment.Comment_ID.Hash = comment_id

	ret, err := rt.db.CommentPhoto(userName, photoID, comment)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error commenting photo")
		w.Write([]byte(ret))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve photo ID from path

	photoID := ps.ByName("photoID")

	// get the user ID

	token := r.Header.Get("user_id")

	userName := ps.ByName("user_name")

	is_valid, err := rt.db.Validate(userName, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	comment_id := ps.ByName("comment_id")

	ret, err := rt.db.UncommentPhoto(userName, photoID, comment_id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error deleting comment")
		w.Write([]byte(ret))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the user ID

	token := r.Header.Get("user_id")

	userName := ps.ByName("user_name")

	is_valid, err := rt.db.Validate(userName, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Retrieve photo from request body

	decoder := json.NewDecoder(r.Body)

	var photo components.Photo

	err = decoder.Decode(&photo)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		w.Write([]byte(components.BadRequestError))

		ctx.Logger.WithError(err).Error("error decoding JSON")

		return // exit the function
	}

	photo_id := ps.ByName("photo_id")

	ret, err := rt.db.UploadPhoto(userName, photo, photo_id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error uploading photo")
		w.Write([]byte(ret))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the user ID

	token := r.Header.Get("user_id")

	userName := ps.ByName("user_name")

	is_valid, err := rt.db.Validate(userName, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	photo_id := ps.ByName("photo_id")

	ret, err := rt.db.DeletePhoto(userName, photo_id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error deleting photo")
		w.Write([]byte(ret))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) getStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the user ID

	token := r.Header.Get("user_id")

	userName := ps.ByName("user_name")

	is_valid, err := rt.db.Validate(userName, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(components.InternalServerError))
		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(components.UnauthorizedError))
		ctx.Logger.Info("unauthorized getstream request received")
		return
	}

	// get bounds from query

	lower_bound_str := r.URL.Query().Get("from")

	if lower_bound_str == "" {
		lower_bound_str = "0"
	}

	lower_bound, err := strconv.Atoi(lower_bound_str)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(components.BadRequestError))
		ctx.Logger.WithError(err).Error("bad getstream request query")
		return

	}

	offset_str := r.URL.Query().Get("offset")

	if offset_str == "" {
		offset_str = "10"
	}

	offset, err := strconv.Atoi(offset_str)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(components.BadRequestError))
		ctx.Logger.WithError(err).Error("bad getstream request query")
		return

	}

	if lower_bound < 0 || offset < 1 || offset > 255 {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(components.BadRequestError))
		ctx.Logger.WithError(err).Error("bad getstream request")
		return

	}

	ret_json_string, err := rt.db.GetStream(userName, lower_bound, offset)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error getting stream")
		w.Write([]byte(ret_json_string))
		return
	}

	w.Write([]byte(ret_json_string))

}
