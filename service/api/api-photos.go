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

	// Retrieve photo ID from path

	photoID := ps.ByName("photo_id")

	// get the photo likes
	ret_data, err := rt.db.GetPhotoLikes(photoID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(ret_data))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error getting photo likes")

		return
	}

	_, err = w.Write([]byte(ret_data))

	if err != nil {
		ctx.Logger.WithError(err).Error("error writing response")
	}

}

func (rt *_router) GetPhotoComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve photo ID from request path

	photoID := ps.ByName("photo_id")

	// get the photo comments
	ret_data, err := rt.db.GetPhotoComments(photoID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, err := w.Write([]byte(ret_data))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error getting photo comments")
		return
	}

	_, err = w.Write([]byte(ret_data))

	if err != nil {
		ctx.Logger.WithError(err).Error("error writing response")
	}

}

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve photo ID from path

	photoID := ps.ByName("photo_id")

	// get the user ID

	token := r.Header.Get("Authorization")
	liker_id := ps.ByName("liker_id")
	liker_name, err := rt.db.GetUsername(liker_id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error getting username")
		return
	}

	is_valid, err := rt.db.Validate(liker_name, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ret, err := rt.db.LikePhoto(liker_id, photoID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error liking photo")
		_, err := w.Write([]byte(ret))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve photo ID from path

	photoID := ps.ByName("photo_id")

	// get the user ID

	token := r.Header.Get("Authorization")

	liker_id := ps.ByName("liker_id")
	liker_name, err := rt.db.GetUsername(liker_id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error getting username")
		return
	}

	is_valid, err := rt.db.Validate(liker_name, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ret, err := rt.db.UnlikePhoto(liker_id, photoID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error unliking photo")
		_, err := w.Write([]byte(ret))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve photo ID from path

	photoID := ps.ByName("photo_id")

	// get the user ID

	token := r.Header.Get("Authorization")

	userName := ps.ByName("user_name")

	// Validate user, get name from header
	commenter_name := r.Header.Get("commenter_name")

	is_valid, err := rt.db.Validate(commenter_name, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, err := w.Write([]byte(components.InternalServerError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte(components.UnauthorizedError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}
		return
	}

	// Retrieve comment from request body

	decoder := json.NewDecoder(r.Body)

	//	tmp_comment := struct {
	//		Comment_ID struct {
	//			Hash string `json:"hash"`
	//		} `json:"comment_id"`
	//		Commenter_Name struct {
	//			Uname string `json:"username-string"`
	//		} `json:"author"`
	//		Comment_Text string `json:"body"`
	//		CreationTime string `json:"creation_time"`
	//		Parent       struct {
	//			Hash string `json:"hash"`
	//		} `json:"parent_post"`
	//	}{}

	comment := components.Comment{}

	err = decoder.Decode(&comment)

	//	comment := components.Comment{
	//		Comment_ID: components.SHA256hash{
	//			Hash: tmp_comment.Comment_ID.Hash,
	//		},
	//		Username: components.User{
	//			Uname: tmp_comment.Commenter_Name.Uname,
	//		},
	//		Body:         tmp_comment.Comment_Text,
	//		CreationTime: time.Time{tmp_comment.CreationTime},
	//		Parent: components.SHA256hash{
	//			Hash: tmp_comment.Parent.Hash,
	//		},
	//	}

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("error decoding JSON")

		_, err := w.Write([]byte(components.BadRequestError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return // exit the function
	}

	comment_id := ps.ByName("comment_id")

	comment.Comment_ID.Hash = comment_id

	ret, err := rt.db.CommentPhoto(userName, photoID, comment)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error commenting photo")
		_, err := w.Write([]byte(ret))
		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Retrieve photo ID from path

	photoID := ps.ByName("photo_id")

	// get the user ID

	token := r.Header.Get("Authorization")

	userName := ps.ByName("user_name")

	// Validate user, get name from header
	commenter_name := r.Header.Get("commenter_name")

	is_valid, err := rt.db.Validate(commenter_name, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, err := w.Write([]byte(components.InternalServerError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)

		_, err := w.Write([]byte(components.UnauthorizedError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return
	}

	comment_id := ps.ByName("comment_id")

	ret, err := rt.db.UncommentPhoto(userName, photoID, comment_id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error deleting comment")
		_, err := w.Write([]byte(ret))
		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the user ID

	token := r.Header.Get("Authorization")
	userName := ps.ByName("user_name")
	is_valid, err := rt.db.Validate(userName, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, err := w.Write([]byte(components.InternalServerError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error validating user")
		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)

		_, err := w.Write([]byte(components.UnauthorizedError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error validating user")

		return
	}

	// Retrieve photo from request body

	decoder := json.NewDecoder(r.Body)
	var photo components.Photo
	err = decoder.Decode(&photo)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		_, err := w.Write([]byte(components.BadRequestError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error decoding JSON")

		return // exit the function
	}

	photo_id := ps.ByName("photo_id")

	ret, err := rt.db.UploadPhoto(userName, photo, photo_id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error uploading photo")
		_, err := w.Write([]byte(ret))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the user ID

	token := r.Header.Get("Authorization")

	userName := ps.ByName("user_name")

	is_valid, err := rt.db.Validate(userName, token)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error validating user")

		_, err := w.Write([]byte(components.InternalServerError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)

		_, err := w.Write([]byte(components.UnauthorizedError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return
	}

	photo_id := ps.ByName("photo_id")

	ret, err := rt.db.DeletePhoto(userName, photo_id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error deleting photo")
		_, err := w.Write([]byte(ret))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (rt *_router) getStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the user ID

	token := r.Header.Get("Authorization")

	userName := ps.ByName("user_name")

	is_valid, err := rt.db.Validate(userName, token)

	if err != nil {

		ctx.Logger.WithError(err).Error("error validating user")

		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(components.InternalServerError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return
	}

	if !is_valid {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte(components.UnauthorizedError))
		ctx.Logger.Info("unauthorized getstream request received")

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return
	}

	// get bounds from query

	lower_bound_str := r.URL.Query().Get("from")

	if lower_bound_str == "" {
		lower_bound_str = "0"
	}

	lower_bound, err := strconv.Atoi(lower_bound_str)

	if err != nil {

		ctx.Logger.WithError(err).Error("bad getstream request query")

		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(components.BadRequestError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return

	}

	offset_str := r.URL.Query().Get("offset")

	if offset_str == "" {
		offset_str = "10"
	}

	offset, err := strconv.Atoi(offset_str)

	if err != nil {

		ctx.Logger.WithError(err).Error("bad getstream request query")

		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(components.BadRequestError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return

	}

	if lower_bound < 0 || offset < 1 || offset > 255 {

		ctx.Logger.WithError(err).Error("bad getstream request")

		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(components.BadRequestError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return

	}

	ret_json_string, err := rt.db.GetStream(userName, lower_bound, offset)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error getting stream")
		_, err := w.Write([]byte(ret_json_string))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		return
	}

	_, err = w.Write([]byte(ret_json_string))

	if err != nil {
		ctx.Logger.WithError(err).Error("error writing response")
	}

}
