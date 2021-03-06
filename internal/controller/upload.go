package controller

import (
	"github.com/0RAJA/Road/internal/logic"
	"github.com/0RAJA/Road/internal/pkg/app"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/bind"
	"github.com/gin-gonic/gin"
)

// Upload 上传服务
// @Summary 上传服务，目前支持图片(png,svg,jpg,webp,bmp,20m以内)和文件(50m以内)
// @Description 上传图片
// @Tags Upload
// @Accept multipart/form-data
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param file formData file true "文件"
// @Param file_type body string true "文件类型 Enums(image,file)"
// @Success 200 {string} string  "上传文件的url"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /upload [post]
func Upload(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	params := logic.UploadParams{}
	valid, errs := bind.BindAndValid(ctx, &params)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParamsErr.WithDetails(bind.FormatBindErr(errs)))
		return
	}
	reply, err := logic.Upload(params)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	response.ToResponse(reply)
}
