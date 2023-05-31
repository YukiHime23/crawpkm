package main

import (
	"github.com/gin-gonic/gin"
	"goCraw/controller"
	"goCraw/database"
	"goCraw/service"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	db, err := database.OpenPostgresDB()
	if err != nil {
		log.Fatal("[main] DB connect error: ", err)
	}

	ser := service.NewAppService(db)

	crawALCtrl := controller.NewCrawALController(ser)
	router.GET("/wallpaper/azur-lane", crawALCtrl.CrawAzurLaneWallpaper)

	crawBookVnCtrl := controller.NewCrawBookVnController(ser)
	router.GET("/book-vn", crawBookVnCtrl.CrawCXB)

	router.GET("/check-heath", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Oke",
		})
	})

	crawPKCtrl := controller.NewCrawPKController(ser)
	router.GET("/craw-data/pokemon-special", crawPKCtrl.CrawPkmSpecial)

	router.Run(":8080")
}
