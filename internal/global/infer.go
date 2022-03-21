package global

import (
	"errors"
	"os"
	"path/filepath"
	"sync"
)

var (
	once    = new(sync.Once)
	RootDir string
)

func exist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || errors.Is(err, os.ErrExist)
}

//计算项目路径
func inferRootDir() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var infer func(path string) string
	infer = func(path string) string {
		if exist(path + "/configs") {
			return path
		}
		return infer(filepath.Dir(path))
	}
	RootDir = infer(cwd)
}

func init() {
	once.Do(func() {
		inferRootDir()
	})
}
