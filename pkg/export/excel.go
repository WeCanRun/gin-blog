package export

import (
	"fmt"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"time"
)

const (
	EXPORT_NAME_FMT = "tags-%s-%s.xlsx"
)

func ExportExcelName(name string, date time.Time) string {
	return fmt.Sprintf(EXPORT_NAME_FMT, name, date.Format("20060102"))
}

func GetExcelSavePath() string {
	return setting.App.ExportSavePath
}

func GetExcelSaveUrl(name string) string {
	return GetExcelSavePath() + name
}

func GetExcelRealDir() string {
	return setting.App.RuntimeRootPath + GetExcelSavePath()
}

func GetExcelRealPath(name string) string {
	return GetExcelRealDir() + name
}

func GetExcelFullUrl(name string) string {
	return setting.App.PrefixUrl + GetExcelSavePath() + name
}
