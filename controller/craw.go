package controller

import (
	"github.com/gin-gonic/gin"
	"goCraw/config"
	"goCraw/model"
	"goCraw/service"
	"net/http"
)

type CrawController struct {
	Pkm service.PokemonSpecialService
	Al  service.AzurLaneService
}

func NewCrawController() *CrawController {
	return &CrawController{
		Pkm: service.NewPokemonSpecialService(),
		Al:  service.NewAzurLaneService(),
	}
}

func (c *CrawController) CrawPkmSpecial(ctx *gin.Context) {
	vol := model.Volume{}
	err := c.Pkm.CrawVolume(config.LinkPkmCraw, vol)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func (c *CrawController) CrawAzurLaneWallpaper(ctx *gin.Context) {

	err := c.Al.CrawWallpaper()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	return
}
