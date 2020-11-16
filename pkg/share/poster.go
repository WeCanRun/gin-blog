package share

import (
	"github.com/WeCanRun/gin-blog/model"
	"github.com/WeCanRun/gin-blog/pkg/file"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

type ArticlePoster struct {
	PosterName string
	*model.Article
	Qr *QrCode
}

func NewArticlePoster(name string, article *model.Article, qrCode *QrCode) *ArticlePoster {
	return &ArticlePoster{
		PosterName: name,
		Article:    article,
		Qr:         qrCode,
	}
}

func GetPosterFlag() string {
	return "poster"
}

// 检验背景图片是否存在
func (ap *ArticlePoster) CheckMergedImage(path string) bool {
	return file.IsExit(path + ap.PosterName)
}

// 获取图片的文件句柄
func (ap *ArticlePoster) OpenMergeImage(path string) (*os.File, error) {
	return file.MustOpen(path, ap.PosterName)
}

type ArticlePosterBg struct {
	Name string
	*ArticlePoster
	*Rect
	*Pt
}

type Rect struct {
	X0 int
	Y0 int
	X1 int
	Y1 int
}

type Pt struct {
	X int
	Y int
}

func NewArticlePosterBg(name string, ap *ArticlePoster, rect *Rect, pt *Pt) *ArticlePosterBg {
	return &ArticlePosterBg{
		Name:          name,
		ArticlePoster: ap,
		Rect:          rect,
		Pt:            pt,
	}
}

// 利用二维码和背景图片生成海报
func (apb *ArticlePosterBg) Generate() (string, string, error) {
	saveDir := GetQrCodeSaveDir()
	// 生成二维码
	path, name, err := apb.Qr.Encode(saveDir)
	if err != nil {
		logging.Error("Generate | Encode fail,err%v", err)
		return path, name, err
	}

	if apb.CheckMergedImage(path) {
		logging.Info("Generate | 海报已存在，不需要合并")
		return path, name, nil
	}

	// 获取合成图片文件句柄
	imageFile, err := apb.OpenMergeImage(path)
	if err != nil {
		logging.Error("Generate | OpenMergeImage fail, err%v", err)
		return path, name, err
	}
	defer imageFile.Close()

	// 获取背景图片文件句柄
	bgFile, err := file.MustOpen(path, apb.Name)
	if err != nil {
		logging.Error("Generate | file.MustOpen fail, err:%v", err)
		return path, name, err
	}
	defer bgFile.Close()

	// 打开二维码
	qrFile, err := file.MustOpen(path, name)
	if err != nil {
		logging.Error("Generate | file.MustOpen fail, err%v", err)
		return path, name, err
	}
	defer qrFile.Close()

	// 解码背景图片
	bgImage, err := jpeg.Decode(bgFile)
	if err != nil {
		logging.Error("Generate | jpeg.Decode fail, err:%v", err)
		return path, name, err
	}
	// 解码二维码
	qrImage, err := jpeg.Decode(qrFile)
	if err != nil {
		logging.Error("Generate | jpeg.Decode fail, err%v", err)
		return path, name, err
	}

	// 开始合成海报
	jpg := image.NewRGBA(image.Rect(apb.Rect.X0, apb.Rect.Y0, apb.Rect.X1, apb.Rect.Y1))
	draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
	draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(apb.Pt.X, apb.Pt.Y)), draw.Over)

	// 将海报写入文件中
	err = jpeg.Encode(imageFile, jpg, nil)
	return path, name, err
}
