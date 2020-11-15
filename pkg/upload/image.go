package upload

import (
	"fmt"
	"github.com/WeCanRun/gin-blog/pkg/file"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/WeCanRun/gin-blog/pkg/util"
	"mime/multipart"
	"os"
	"strings"
)

func GetImageFullUrl(name string) string {
	return GetImagePrefix() + GetImagePath() + name
}

func GetImagePrefix() string {
	return setting.App.PrefixUrl
}

func GetImagePath() string {
	return setting.App.ImageSavePath
}

func GetImageName(name string) string {
	ext := file.GetExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

func GetImageFullPath() string {
	return setting.App.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(name string) bool {
	ext := file.GetExt(name)
	for _, allowExt := range setting.App.ImageAllowExts {
		if strings.ToUpper(ext) == strings.ToUpper(allowExt) {
			return true
		}
	}
	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		logging.Error("file.GetSize fail,%v", err)
		return false
	}
	return size <= int(setting.App.ImageMaxSize)
}

func CheckImage(src string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd() pwd:%s, err: %v", pwd, err)
	}
	dir := pwd + "/" + src
	if err := file.IsNotExitMKDir(dir); err != nil {
		return fmt.Errorf("file.IsNotExitMKDir dir: %s, err: %v", dir, err)
	}

	if perm := file.NotPermission(src); perm {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	return nil
}
