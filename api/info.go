package api

import (
	"infoweb/service"

	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InfoList(c *gin.Context) {
	service := service.InfoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetInfoList()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}