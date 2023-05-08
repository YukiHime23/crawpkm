package main

import (
	"github.com/gin-gonic/gin"
	"goCraw/controller"
	"net/http"
)

func main() {
	router := gin.Default()

	crawCtrl := controller.NewCrawController()
	router.GET("/check-heath", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Oke",
		})
	})
	router.GET("/craw-data/pokemon-special", crawCtrl.CrawPkmSpecial)
	router.GET("/wallpaper/azur-lane", crawCtrl.CrawAzurLaneWallpaper)

	router.Run(":8080")
}
