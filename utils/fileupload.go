package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

//FileUpload FileUpload
func FileUpload(file multipart.File, filename string, destinationDir string) error {

	if _, err := os.Stat(destinationDir); os.IsNotExist(err) {
		os.Mkdir(destinationDir, 0700)
	}
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", destinationDir, filename), os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return err
	}
	return nil

}
