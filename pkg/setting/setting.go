package setting

import (
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"log"
	"time"
)
import "github.com/go-ini/ini"

const (
	base     = "base"
	server   = "server"
	app      = "app"
	database = "database"
)

var (
	// [base]
	Cfg     *ini.File
	RunMode string
	// [server]
	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	// [app]
	PageSize  uint
	JwtSecret string
	// [database]
	DbType, DbName, User, Password, Host, TablePrefix string
)

func init() {
	var err error
	Cfg, err = ini.Load("./conf/app.ini")
	if err != nil {
		logging.Fatal("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadDataBase()
	LoadServer()
	LoadApp()
}

func LoadDataBase() {
	sec, err := Cfg.GetSection(database)
	if err != nil {
		log.Fatalf("fail to get setion %s: %v", database, err)
	}
	DbType = sec.Key("TYPE").String()
	DbName = sec.Key("NAME").String()
	User = sec.Key("USER").String()
	Password = sec.Key("PASSWORD").String()
	Host = sec.Key("HOST").String()
	TablePrefix = sec.Key("TABLE_PREFIX").String()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("fail to get setion 'server': %v", err)
	}
	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeOut = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("fail to get setion 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = uint(sec.Key("PAGE_SIZE").MustInt(10))
}
