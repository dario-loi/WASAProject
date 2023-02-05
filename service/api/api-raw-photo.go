package api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
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

	const well_known = "default"

	// Get photo id from path

	uuid := ps.ByName("UUID")

	if uuid == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(components.BadRequestError))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.Error("error getting photo id")
		return
	}

	// Check if the picture exists
	if uuid != well_known {
		exists, err := rt.db.CheckPhotoExists(uuid)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte(fmt.Errorf(components.NotFoundErrorF, err).Error()))

			if err != nil {
				ctx.Logger.WithError(err).Error("error writing response")
			}

			ctx.Logger.WithError(err).Error(fmt.Sprintf("error checking photo (uuid: %s) existence, details: %s", uuid, err))
			return
		}
		if !exists {
			// If the photo doesn't exist, return the default one
			uuid = well_known
		}
	}

	// Get the photo from the filesystem

	img_file, err := os.Open("photos/" + uuid + ".png")

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(fmt.Errorf(components.InternalServerErrorF, err).Error()))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error opening photo")
		return
	}

	defer func() {
		_ = img_file.Close()
	}()

	img, _, err := image.Decode(img_file)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(fmt.Errorf(components.InternalServerErrorF, err).Error()))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error decoding photo")
		return
	}

	// Encode the photo as PNG

	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(fmt.Errorf(components.InternalServerErrorF, err).Error()))

		if err != nil {
			ctx.Logger.WithError(err).Error("error writing response")
		}

		ctx.Logger.WithError(err).Error("error encoding photo")
		return
	}

	bin := buf.Bytes()

	w.Write([]byte("data:image/png;base64,"))

	// Send the photo to the client
	_, err = w.Write([]byte(toBase64(bin)))

	if err != nil {
		ctx.Logger.WithError(err).Error("error writing response")
	}

}
