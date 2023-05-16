package controller

import (
	"github.com/gin-gonic/gin"
	"goCraw/config"
	"goCraw/model"
	"goCraw/service"
	"net/http"
)

type CrawPKController struct {
	service service.PokemonSpecialService
}

func NewCrawPKController(s service.PokemonSpecialService) *CrawPKController {
	return &CrawPKController{service: s}
}

func (c *CrawPKController) CrawPkmSpecial(ctx *gin.Context) {
	vol := model.Volume{}
	err := c.service.CrawVolume(config.LinkPkmCraw, vol)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}
