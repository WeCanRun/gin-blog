package file

import (
	"io/ioutil"
	"log"
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

// 打开一个文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

// 成功打开文件
func MustOpen(path, name string) (*os.File, error) {
	err := IsNotExitMKDir(path)
	if err != nil {
		log.Printf("MustOpen | IsNotExitMKDir fail, 文件打开失败 err:%v\n", err)
		return nil, err
	}
	file, err := Open(path+name, os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Printf("MustOpen | Open 文件打开失败,err:%v\n", err)
		return file, err
	}
	return file, nil
}
