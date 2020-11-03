package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}

func GetExt(fileName string) string {
	return path.Ext(fileName)
}

func IsExit(path string) bool {
	_, err := os.Stat(path)
	return os.IsExist(err)
}

func NotPermission(path string) bool {
	_, err := os.Stat(path)
	return os.IsPermission(err)
}

func IsNotExitMKDir(path string) error {
	// 目录存在或者创建成功返回 nil
	if !IsExit(path) {
		return MKDir(path)
	}
	return nil
}

func MKDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}
