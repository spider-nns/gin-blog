package app

import (
	"github.com/gin-gonic/gin"
	"http-gin/pkg/errcode"
	"net/http"
)

//响应处理

type Response struct {
	Ctx *gin.Context
}
type Pager struct {
	PageNo    int `json:"pageNo"`
	PageSize  int `json:"pageSize"`
	TotalRows int `json:"totalRows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			PageNo:    GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
