package service

import (
	"http-gin/internal/model"
	"http-gin/pkg/app"
)

//入参校验
type CountTagRequest struct {
	Name  string `form:"name" binding:"max=20"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}
type TagListRequest struct {
	Name  string `form:"name" binding:"max=20"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}
type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,gte=1"`
	CreatedBy string `form:"createdBy" binding:"required,min=2,max=20"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}
type UpdateTagRequest struct {
	ID         uint32 `form:"id" json:"id"binding:"required,gte=1"`
	Name       string `form:"name" json:"name" binding:"required,min=2,max=20"`
	State      uint8  `form:"state" json:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modifiedBy" json:"modified_by" binding:"required,min=2,max=20"`
}
type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}
func (svc *Service) GetTagList(param *TagListRequest, paper *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, paper.PageNo, paper.PageSize)
}
func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}
func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}
func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
