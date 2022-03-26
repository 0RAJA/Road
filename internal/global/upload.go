package global

import "github.com/0RAJA/Road/internal/pkg/upload"

var Upload *upload.Upload

func init() {
	image := upload.NewType(upload.FileType(AllSetting.Upload.Image.Type), AllSetting.Upload.Image.Suffix, AllSetting.Upload.Image.MaxSize, AllSetting.Upload.Image.UrlPrefix, AllSetting.Upload.Image.LocalPath)
	file := upload.NewType(upload.FileType(AllSetting.Upload.File.Type), AllSetting.Upload.File.Suffix, AllSetting.Upload.File.MaxSize, AllSetting.Upload.File.UrlPrefix, AllSetting.Upload.File.LocalPath)
	Upload = upload.Init(image, file)
}
