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
		fmt.Println("Empty photo id")
		return
	}

	// Check if the picture exists

	exists, err := rt.db.CheckPhotoExists(uuid)

	if err != nil || !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Errorf(components.NotFoundErrorF, err).Error()))
		fmt.Println("Requested photo does not exist")
		return
	}

	// Get the photo from the filesystem

	img_file, err := os.Open("photos/" + uuid + ".png")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf(components.InternalServerErrorF, err).Error()))
		fmt.Println(fmt.Errorf("error opening photo: %w", err))
		return
	}

	defer func() {
		_ = img_file.Close()
	}()

	img, _, err := image.Decode(img_file)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf(components.InternalServerErrorF, err).Error()))
		fmt.Println(fmt.Errorf("error decoding photo: %w", err))
		return
	}

	// Encode the photo as JPEG

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf(components.InternalServerErrorF, err).Error()))
		fmt.Println(fmt.Errorf("error encoding photo: %w", err))
		return
	}

	bin := buf.Bytes()

	// Send the photo to the client
	w.Write([]byte(toBase64(bin)))

}
