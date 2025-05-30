package crawpkm

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func generateFileName(url, nameReplace string) string {
	ext := ".jpg"
	fileName := path.Base(url)
	if path.Ext(fileName) != "" {
		return fileName
	}
	fmt.Println(nameReplace, ">>>>> file name", fileName)

	return nameReplace + ext
}

func DownloadFile(url, nameReplace, pathTo string) error {
	fileName := generateFileName(url, nameReplace)

	//Get the response bytes from the url
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}

	file, err := os.Create(pathTo + "/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the field
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Println("-> download done \"" + fileName + "\" <-")
	return nil
}
