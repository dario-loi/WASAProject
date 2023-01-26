package api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"net/http"
	"os"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/components"
	"github.com/julienschmidt/httprouter"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get photo id from path

	uuid := ps.ByName("UUID")

	if uuid == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(components.BadRequestError))
		ctx.Logger.Error("error getting photo id")
		return
	}

	// Check if the picture exists

	exists, err := rt.db.CheckPhotoExists(uuid)

	if err != nil || !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Errorf(components.NotFoundErrorF, err).Error()))
		ctx.Logger.WithError(err).Error("error checking photo existence")
		return
	}

	// Get the photo from the filesystem

	img_file, err := os.Open("photos/" + uuid + ".png")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf(components.InternalServerErrorF, err).Error()))
		ctx.Logger.WithError(err).Error("error opening photo")
		return
	}

	defer func() {
		_ = img_file.Close()
	}()

	img, _, err := image.Decode(img_file)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf(components.InternalServerErrorF, err).Error()))
		ctx.Logger.WithError(err).Error("error decoding photo")
		return
	}

	// Encode the photo as JPEG

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf(components.InternalServerErrorF, err).Error()))
		ctx.Logger.WithError(err).Error("error encoding photo")
		return
	}

	bin := buf.Bytes()

	// Send the photo to the client
	w.Write([]byte(toBase64(bin)))

}
