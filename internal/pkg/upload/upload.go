package upload

import (
	"github.com/0RAJA/Road/internal/pkg/app/errcode"
	"mime/multipart"
	"os"
)

func (u *Upload) SaveFile(fileType FileType, fileHeader *multipart.FileHeader) (string, error) {
	fileName, ext := GetFileName(fileHeader.Filename) //获取加密文件名
	fileTypeor, err := checkContainExt(fileType, ext)
	if err != nil { //判断文件类型是否合法
		return "", err
	}
	if !checkMaxSize(fileTypeor, fileHeader) { //检查文件大小
		return "", errcode.FileSizeErr
	}
	uploadSavePath := fileTypeor.GetPath()
	if CheckSavePath(uploadSavePath) { //检查保存路径是否存在
		if err := createSavePath(uploadSavePath, os.ModePerm); err != nil { //创建保存路径
			return "", errcode.CreatePathErr
		}
	}
	if checkPermission(uploadSavePath) { //检查权限
		return "", errcode.CompetenceErr
	}
	dst := uploadSavePath + "/" + fileName + ext //加密文件名
	if err := saveFile(fileHeader, dst); err != nil {
		return "", err
	}
	accessUrl := fileTypeor.GetUrlPrefix() + "/" + fileName + ext
	return accessUrl, nil
}
