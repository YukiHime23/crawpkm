package service

import (
	"encoding/json"
	"goCraw/config"
	"goCraw/model"
	"io"
	"net/http"
	"os"
	"strings"
)

type AzurLaneService interface {
	CrawWallpaper() error
}

func (a AppService) CrawWallpaper() error {
	res, err := http.Get(config.ApiListWallpaperAzurLane)
	if err != nil {
		return err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var resApi model.ResponseApi
	if err = json.Unmarshal(resBody, &resApi); err != nil {
		return err
	}

	pathFile := "AzurLane"
	if err = os.MkdirAll(pathFile, os.ModePerm); err != nil {
		return err
	}

	var list []model.AzurLane
	var idExist []int

	// get id exist
	a.db.Select("id_wallpaper").Table("azur_lanes").Scan(&idExist)

	for _, row := range resApi.Data.Rows {
		if IntInArray(idExist, row.ID) {
			continue
		}

		var al model.AzurLane
		al.Url = config.DomainLoadWallpaperAzurLane + row.Works
		al.FileName = strings.ReplaceAll(row.Title+" ("+row.Artist+").jpeg", "/", "-")
		al.IdWallpaper = row.ID
		if err = DownloadFile(al.Url, al.FileName, pathFile); err != nil {
			return err
		}
		list = append(list, al)
	}
	if len(list) > 0 {
		a.db.Create(&list)
	}
	return nil
}
