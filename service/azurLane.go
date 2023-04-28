package service

import (
	"encoding/json"
	"fmt"
	"goCraw/config"
	"goCraw/domain"
	"io"
	"net/http"
	"os"
	"strings"
)

type AzurLaneService interface {
	CrawWallpaper() error
}

type azurLaneService struct {
}

func (a azurLaneService) CrawWallpaper() error {
	res, err := http.Get(config.ApiListWallpaperAzurLane)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var resApi domain.ResponseApi
	if err = json.Unmarshal([]byte(resBody), &resApi); err != nil {
		return err
	}

	pathFile := "AzurLane"
	if err = os.MkdirAll(pathFile, os.ModePerm); err != nil {
		return err
	}

	// 95 file
	for _, row := range resApi.Data.Rows {
		fmt.Println("-> Start download <-")
		urlWall := config.DomainLoadWallpaperAzurLane + row.Works
		fileName := strings.ReplaceAll(row.Title+" ("+row.Artist+").jpeg", "/", "-")
		if err = DownloadFile(urlWall, fileName, pathFile); err != nil {
			return err
		}
		fmt.Println("-> download done \"" + fileName + "\" <-")
	}
	return nil
}

func NewAzurLaneService() AzurLaneService {
	return &azurLaneService{}
}
