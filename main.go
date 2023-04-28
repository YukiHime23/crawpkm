package main

import (
	"github.com/gin-gonic/gin"
	"goCraw/controller"
)

func main() {
	router := gin.Default()

	crawCtrl := controller.NewCrawController()
	router.GET("/craw-data/pokemon-special", crawCtrl.CrawPkmSpecial)
	router.GET("/wallpaper/azur-lane", crawCtrl.CrawAzurLaneWallpaper)

	router.Run(":8181")
}
