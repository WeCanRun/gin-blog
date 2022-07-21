package share

import (
	"github.com/WeCanRun/gin-blog/internal/model"
	"github.com/WeCanRun/gin-blog/pkg/file"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/golang/freetype"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
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
	bgFile, err := file.MustOpen(setting.APP.BgSavePath, apb.Name)
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

	// 生成海报并写入文字
	err = apb.WriterFontToPoster(&DrawText{
		JPG:     jpg,
		Merged:  imageFile,
		Title:   "这是海报大标题",
		X0:      80,
		Y0:      160,
		Size0:   42,
		SubText: "—无意",
		X1:      320,
		Y1:      220,
		size1:   36,
	}, "msyhbd.ttc")

	return path, name, err
}

type DrawText struct {
	JPG    draw.Image
	Merged *os.File

	Title string
	X0    int
	Y0    int
	Size0 float64

	SubText string
	X1      int
	Y1      int
	size1   float64
}

func (apb *ArticlePosterBg) WriterFontToPoster(d *DrawText, fontName string) error {
	// 打开字体文件
	fontSrc := setting.APP.FontSavePath + fontName
	fontSrcBytes, err := ioutil.ReadFile(fontSrc)
	if err != nil {
		logging.Error("DrawPoster | ioutil.ReadFile fail, fontSrc:%s, err:%v", fontSrc, err)
		return err
	}

	// 解析字体
	font, err := freetype.ParseFont(fontSrcBytes)
	if err != nil {
		logging.Error("DrawPoster | freetype.ParseFont 解析字体失败, err:%v", err)
		return err
	}

	// 设置字体属性
	ftContext := freetype.NewContext()
	ftContext.SetDPI(72)
	ftContext.SetFont(font)
	ftContext.SetFontSize(d.Size0)
	ftContext.SetClip(d.JPG.Bounds())
	ftContext.SetDst(d.JPG)
	ftContext.SetSrc(image.Black)

	// 写入主标题
	pt0 := freetype.Pt(d.X0, d.Y0)
	_, err = ftContext.DrawString(d.Title, pt0)
	if err != nil {
		logging.Error("DrawPoster | DrawString fail, err:%v", err)
		return err
	}

	// 写入副标题
	ftContext.SetFontSize(d.size1)
	pt1 := freetype.Pt(d.X1, d.Y1)
	_, err = ftContext.DrawString(d.SubText, pt1)
	if err != nil {
		logging.Error("DrawPoster | DrawString fail, err:%v", err)
		return err
	}

	// 写入图片中
	err = jpeg.Encode(d.Merged, d.JPG, nil)
	if err != nil {
		logging.Error("DrawPoster | Encode fail,err:%v", err)
	}
	return err
}
