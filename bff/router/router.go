package router

import (
	"day6/bff/handler/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	cms := r.Group("/cms")
	{
		cms.POST("/create", service.CmsCreate)
		cms.POST("/del", service.CmsDel)
		cms.POST("/update", service.CmsUpdate)
		cms.GET("/list", service.CmsList)
	}
	return r
}
