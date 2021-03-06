package app

import (
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

//εεΊε€η

type Response struct {
	Ctx *gin.Context
}

type ResponseList struct {
	List  interface{} `json:"list"`
	Pager Pager       `json:"page"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, ResponseList{
		List: list,
		Pager: Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	r.Ctx.JSON(err.StatusCode(), err)
}
