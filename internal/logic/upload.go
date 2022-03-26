package logic

import (
	"github.com/0RAJA/Road/internal/global"
	"github.com/0RAJA/Road/internal/pkg/upload"
)

func Upload(params UploadParams) (string, error) {
	url, err := global.Upload.SaveFile(upload.FileType(params.FileType), params.File)
	if err != nil {
		global.Logger.Info(err.Error())
		return "", err
	}
	return url, nil
}
