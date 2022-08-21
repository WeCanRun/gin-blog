package export

import (
	"fmt"
	"github.com/WeCanRun/gin-blog/global"
	"time"
)

const (
	EXPORT_NAME_FMT = "tags-%s-%s.xlsx"
)

func ExportExcelName(name string, date time.Time) string {
	return fmt.Sprintf(EXPORT_NAME_FMT, name, date.Format("20060102"))
}

func GetExcelSavePath() string {
	return global.Setting.APP.ExportSavePath
}

func GetExcelSaveUrl(name string) string {
	return GetExcelSavePath() + name
}

func GetExcelRealDir() string {
	return global.Setting.APP.RuntimeRootPath + GetExcelSavePath()
}

func GetExcelRealPath(name string) string {
	return GetExcelRealDir() + name
}

func GetExcelFullUrl(name string) string {
	return global.Setting.APP.PrefixUrl + GetExcelSavePath() + name
}
