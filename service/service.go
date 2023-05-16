package service

import (
	"errors"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"path"
)

type AppService struct {
	db *gorm.DB
}

func NewAppService(db *gorm.DB) *AppService {
	return &AppService{db: db}
}

func DownloadFile(URL, fileName string, pathTo string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}

	//Create an empty file
	if fileName == "" {
		fileName = path.Base(URL)
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

	return nil
}
