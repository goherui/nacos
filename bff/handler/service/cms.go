package service

import (
	"day6/bff/basic/config"
	"day6/bff/handler/request"
	"day6/bff/handler/response"
	__ "day6/proto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CmsCreate(c *gin.Context) {
	var form request.CmsCreate
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "参数错误",
		})
		return
	}
	r, err := config.CmsClient.CmsCreate(c, &__.CmsCreateReq{
		Title:      form.Title,
		CategoryId: int64(form.CategoryID),
		Content:    form.Content,
		Status:     int64(form.Status),
		ViewCount:  int64(form.ViewCount),
		Creator:    int64(form.Creator),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusOK, response.CmsCreate{
		Code: int(r.Code),
		Msg:  r.Msg,
	})
}
func CmsDel(c *gin.Context) {
	var form request.CmsDel
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "参数错误",
		})
		return
	}
	r, err := config.CmsClient.CmsDel(c, &__.CmsDelReq{
		Id: int64(form.Id),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusOK, response.CmsDel{
		Code: int(r.Code),
		Msg:  r.Msg,
	})
}
func CmsUpdate(c *gin.Context) {
	var form request.CmsUpdate
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "参数错误",
		})
		return
	}
	r, err := config.CmsClient.CmsUpdate(c, &__.CmsUpdateReq{
		Id:         int64(form.Id),
		Title:      form.Title,
		CategoryId: int64(form.CategoryID),
		Content:    form.Content,
		Status:     int64(form.Status),
		ViewCount:  int64(form.ViewCount),
		Creator:    int64(form.Creator),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusOK, response.CmsUpdate{
		Code: int(r.Code),
		Msg:  r.Msg,
	})
}
func CmsList(c *gin.Context) {
	var form request.CmsList
	// 根据 Content-Type Header 推断使用哪个绑定器。
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   "参数错误",
		})
		return
	}
	r, err := config.CmsClient.CmsList(c, &__.CmsListReq{
		Page: int64(form.Page),
		Size: int64(form.Size),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	var list []response.Cms
	for _, cms := range r.List {
		list = append(list, response.Cms{
			Id:           int(cms.Id),
			Title:        cms.Title,
			CategoryName: cms.CategoryName,
			Content:      cms.Content,
			Status:       int8(cms.Status),
			ViewCount:    uint(cms.ViewCount),
			Creator:      uint64(cms.Creator),
		})
	}
	c.JSON(http.StatusOK, response.CmsList{
		List: list,
		Code: int(r.Code),
		Msg:  r.Msg,
	})
}
