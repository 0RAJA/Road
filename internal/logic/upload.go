package logic

import (
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"github.com/0RAJA/Road/internal/pkg/upload"
)

func Upload(params UploadParams) (string, *errcode.Error) {
	url, err := global.Upload.SaveFile(upload.FileType(params.FileType), params.File)
	if err != nil {
		global.Logger.Info(err.Error())
		return "", errcode.ServerErr
	}
	return url, nil
}
