package controller

import (
	"github.com/gin-gonic/gin"
	"goCraw/service"
	"net/http"
)

type CrawBookVnController struct {
	service service.BookVnService
}

func NewCrawBookVnController(s service.BookVnService) *CrawBookVnController {
	return &CrawBookVnController{service: s}
}

func (c *CrawBookVnController) CrawCXB(ctx *gin.Context) {
	res := c.service.CrawBookVn()

	ctx.JSON(http.StatusOK, gin.H{
		"message": http.StatusText(http.StatusOK),
		"data":    res,
	})
	return
}
