package share

import (
	"github.com/WeCanRun/gin-blog/pkg/file"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/WeCanRun/gin-blog/pkg/util"
	"github.com/boombuler/barcode"
	_ "github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/jpeg"
	"os"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const EXT_JPG = ".jpg"

func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Ext:    EXT_JPG,
		Level:  level,
		Mode:   mode,
	}
}

func GetQrCodePath() string {
	return setting.APP.QrCodeSavePath
}

func GetQrCodeSaveDir() string {
	return setting.APP.RuntimeRootPath + GetQrCodePath()
}

func GetQrCodeSavePath(name string) string {
	return GetQrCodeSaveDir() + name
}

// 获取文件名
func GetQrCodeFileName(url string) string {
	return util.EncodeMD5(url)
}

func GetQrCodeFullUrl(name string) string {
	return setting.APP.PrefixUrl + GetQrCodePath() + name
}

// 获取二维码的文件后缀
func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

// 检验是否能进行编码
func (q *QrCode) CheckEncode(dir string) bool {
	src := dir + GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	return file.IsExit(src)
}

// 进行二维码编码
func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	fullName := path + name
	if file.IsExit(fullName) {
		logging.Info("Encode | 二维码 %s 已存在，无需编码", fullName)
		return path, name, nil
	}
	// 编码
	encode, err := qr.Encode(fullName, q.Level, q.Mode)
	if err != nil {
		logging.Error("Encode | qr.Encode fail, err:%v", err)
		return path, name, err
	}
	// 重置二维码尺寸
	encode, err = barcode.Scale(encode, q.Width, q.Height)
	if err != nil {
		logging.Error("Encode | qr.Encode fail, err:%v", err)
		return path, name, err
	}
	// 获取保存二维码的文件句柄
	err = file.IsNotExitMKDir(path)
	if err != nil {
		logging.Error("Encode#file.Open | 目录不存在，创建目录失败, err:%v", err)
		return path, name, err
	}
	f, err := file.Open(fullName, os.O_CREATE, os.ModePerm)
	defer f.Close()
	if err != nil {
		logging.Error("Encode#file.Open | 打开文件失败, err:%v", err)
		return path, name, err
	}
	// 将二维码写入文件中
	err = jpeg.Encode(f, encode, nil)
	if err != nil {
		logging.Error("Encode#jpeg.Encode | 二维码写入文件失败, err:%v", err)
		return path, name, err
	}

	return path, name, nil
}
