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
	ImagePrefixUrl  string   `json:"image_prefix_url"`
	ImageSavePath   string   `json:"image_save_path"`
	ImageMaxSize    uint     `json:"image_max_size"`
	ImageAllowExts  []string `json:"image_allow_exts"`
	LogSavePath     string   `json:"log_save_path"`
	LogSaveName     string   `json:"log_save_name"`
	LogFileExt      string   `json:"log_file_ext"`
	TimeFormat      string   `json:"time_format"`
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

const (
	SERVER   = "server"
	APP      = "app"
	DATABASE = "database"
)

var (
	Database = &database{}
	Server   = &server{}
	App      = &app{}
	Cfg      *ini.File
)

func Setup() {
	var err error
	Cfg, err = ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadDataBase()
	LoadServer()
	LoadApp()
}

func LoadDataBase() {
	err := Cfg.Section(DATABASE).MapTo(Database)
	if err != nil {
		log.Fatalf("Cfg.MapTo App fail: %v, %v", DATABASE, err)
	}
}

func LoadServer() {
	err := Cfg.Section(SERVER).MapTo(Server)
	if err != nil {
		log.Fatalf("Cfg.MapTo App fail: %v, %v", SERVER, err)
	}
	Server.ReadTimeout *= time.Second
	Server.WriteTimeout *= time.Second
}

func LoadApp() {
	err := Cfg.Section(APP).MapTo(App)
	if err != nil {
		log.Fatalf("Cfg.MapTo App fail: %v,%v", APP, err)
	}

	App.ImageMaxSize *= 1024 * 1024
}
