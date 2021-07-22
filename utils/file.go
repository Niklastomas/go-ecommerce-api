package utils

import (
	"io"
	"net/http"
	"os"
)

func UploadFile(r *http.Request) (string, error) {
	// Max upload size == 10 MB
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	dst, err := os.Create("./media/" + header.Filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		return "", err
	}
	return header.Filename, nil

}
