package service

import (
	"context"
	__ "day6/proto"
	"day6/service/basic/config"
	"day6/service/model"
	"net/http"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedStreamGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) CmsCreate(_ context.Context, in *__.CmsCreateReq) (*__.CmsCreateResp, error) {
	var content model.CmsContent
	err := content.FindTitle(config.DB, in.Title)
	if err == nil {
		return &__.CmsCreateResp{
			Code: http.StatusBadRequest,
			Msg:  "内容已存在",
		}, nil
	}
	content = model.CmsContent{
		Title:      in.Title,
		CategoryID: uint64(in.CategoryId),
		Content:    in.Content,
		Status:     int8(in.Status),
		ViewCount:  uint(in.ViewCount),
		Creator:    uint64(in.Creator),
	}
	err = content.CmsCreate(config.DB)
	if err != nil {
		return &__.CmsCreateResp{
			Code: http.StatusBadRequest,
			Msg:  "添加失败",
		}, nil
	}
	return &__.CmsCreateResp{
		Code: http.StatusOK,
		Msg:  "添加成功",
	}, nil
}
func (s *Server) CmsDel(_ context.Context, in *__.CmsDelReq) (*__.CmsDelResp, error) {
	var content model.CmsContent
	err := content.FindId(config.DB, in.Id)
	if err != nil {
		return &__.CmsDelResp{
			Code: http.StatusBadRequest,
			Msg:  "内容不存在",
		}, nil
	}
	err = content.CmsDel(config.DB, in.Id)
	if err != nil {
		return &__.CmsDelResp{
			Code: http.StatusBadRequest,
			Msg:  "删除失败",
		}, nil
	}
	return &__.CmsDelResp{
		Code: http.StatusOK,
		Msg:  "删除成功",
	}, nil
}
func (s *Server) CmsUpdate(_ context.Context, in *__.CmsUpdateReq) (*__.CmsUpdateResp, error) {
	var content model.CmsContent
	err := content.FindTitle(config.DB, in.Title)
	if err != nil {
		return &__.CmsUpdateResp{
			Code: http.StatusBadRequest,
			Msg:  "内容不存在",
		}, nil
	}
	content = model.CmsContent{
		Title:      in.Title,
		CategoryID: uint64(in.CategoryId),
		Content:    in.Content,
		Status:     int8(in.Status),
		ViewCount:  uint(in.ViewCount),
		Creator:    uint64(in.Creator),
	}
	err = content.CmsUpdate(config.DB, in.Id)
	if err != nil {
		return &__.CmsUpdateResp{
			Code: http.StatusBadRequest,
			Msg:  "修改失败",
		}, nil
	}
	return &__.CmsUpdateResp{
		Code: http.StatusOK,
		Msg:  "修改成功",
	}, nil
}
func (s *Server) CmsList(_ context.Context, in *__.CmsListReq) (*__.CmsListResp, error) {
	var content model.CmsContent
	list, err := content.FindContent(config.DB, in)
	if err != nil {
		return &__.CmsListResp{
			Code: http.StatusBadRequest,
			Msg:  "获取列表失败",
		}, nil
	}
	var lists []*__.Cms

	for _, cmsContent := range list {
		lists = append(lists, &__.Cms{
			Id:           cmsContent.Id,
			Title:        cmsContent.Title,
			CategoryName: cmsContent.CategoryName,
			Content:      cmsContent.Content,
			Status:       cmsContent.Status,
			ViewCount:    cmsContent.ViewCount,
			Creator:      cmsContent.Creator,
		})
	}
	return &__.CmsListResp{
		List: lists,
		Code: http.StatusOK,
		Msg:  "列表获取成功",
	}, nil
}
