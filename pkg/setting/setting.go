package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type app struct {
	JwtSecret       string   `json:"jwt_secret"`
	PageSize        uint     `json:"page_size"`
	RuntimeRootPath string   `json:"runtime_root_path"`
	PrefixUrl       string   `json:"prefix_url"`
	ImageSavePath   string   `json:"image_save_path"`
	ImageMaxSize    uint     `json:"image_max_size"`
	ImageAllowExts  []string `json:"image_allow_exts"`
	LogSavePath     string   `json:"log_save_path"`
	LogSaveName     string   `json:"log_save_name"`
	LogFileExt      string   `json:"log_file_ext"`
	TimeFormat      string   `json:"time_format"`
	ExportSavePath  string   `json:"export_save_path"`
	QrCodeSavePath  string   `json:"qr_code_save_path"`
	FontSavePath    string   `json:"font_save_path"`
	BgSavePath      string   `json:"bg_save_path"`
}

type server struct {
	RunMode      string        `json:"run_mode"`
	HttpPort     int           `json:"http_port"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
}

type database struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	DbName      string `json:"db_name"`
	TablePrefix string `json:"table_prefix"`
}

type redis struct {
	Host        string        `json:"host"`
	Password    string        `json:"password"`
	MaxIdle     int           `json:"max_idle"`
	MaxActive   int           `json:"max_active"`
	IdleTimeout time.Duration `json:"idle_timeout"`
}

const (
	SERVER   = "server"
	APP      = "app"
	DATABASE = "database"
	REDIS    = "redis"
)

var (
	Redis    = &redis{}
	Database = &database{}
	Server   = &server{}
	App      = &app{}
	Cfg      *ini.File
)

func Setup(path string) {
	var err error
	Cfg, err = ini.Load(path)
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	loadDataBase()
	loadServer()
	loadApp()
	loadRedis()
}

func loadRedis() {
	err := Cfg.Section(REDIS).MapTo(Redis)
	if err != nil {
		log.Fatalf("Cfg.MapTo Redis fail: %v, %v", REDIS, err)
	}
	Redis.IdleTimeout *= time.Second
}

func loadDataBase() {
	err := Cfg.Section(DATABASE).MapTo(Database)
	if err != nil {
		log.Fatalf("Cfg.MapTo Database fail: %v, %v", DATABASE, err)
	}
}

func loadServer() {
	err := Cfg.Section(SERVER).MapTo(Server)
	if err != nil {
		log.Fatalf("Cfg.MapTo Server fail: %v, %v", SERVER, err)
	}
	Server.ReadTimeout *= time.Second
	Server.WriteTimeout *= time.Second
}

func loadApp() {
	err := Cfg.Section(APP).MapTo(App)
	if err != nil {
		log.Fatalf("Cfg.MapTo App fail: %v,%v", APP, err)
	}

	App.ImageMaxSize *= 1024 * 1024
}
