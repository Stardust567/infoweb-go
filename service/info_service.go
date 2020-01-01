package service

import (
	"infoweb/model"
	"infoweb/serializer"
)

// ListVideoService 视频列表服务
type InfoService struct {
	Title string `form:"title"`
	City  string `form:"city"`
	Limit int `form:"limit"`
}

// List 视频列表
func (service *InfoService) GetInfoList() serializer.Response {
	InfoList := []model.Info{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 6
	}

	service.Title = "%" + service.Title + "%"
	service.City = "%" + service.City + "%"

	if err := model.DB.Model(model.Info{}).Where("title LIKE ? AND city LIKE ?", service.Title, 
	service.City).Count(&total).Error; err != nil {
		return serializer.Response{
			Code: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Where("title LIKE ? AND city LIKE ?", service.Title, 
	service.City).Limit(service.Limit).Order("time desc").Find(&InfoList).Error; err != nil {
		return serializer.Response{
			Code: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildInfoList(InfoList), uint(total))
}