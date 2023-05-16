package controller

import (
	"github.com/gin-gonic/gin"
	"goCraw/service"
	"net/http"
)

type CrawALController struct {
	service service.AzurLaneService
}

func NewCrawALController(s service.AzurLaneService) *CrawALController {
	return &CrawALController{service: s}
}

func (c *CrawALController) CrawAzurLaneWallpaper(ctx *gin.Context) {
	err := c.service.CrawWallpaper()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	return
}
