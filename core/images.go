package core

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
)

func SaveImage(id string, b string) (error, string) {

	url := fmt.Sprintf("storage/user-profile-image-%v.png", id)

	unbased, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return err, url
	}

	r := bytes.NewReader(unbased)

	im, err := png.Decode(r)
	if err != nil {
		im, err = jpeg.Decode(r)
		if err != nil {
			return err, url
		}
		f, err := os.OpenFile(url, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return err, url
		}

		jpeg.Encode(f, im, &jpeg.Options{Quality: 80})

		return err, url
	}

	f, err := os.OpenFile(url, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err, url
	}

	png.Encode(f, im)

	return err, url
}
