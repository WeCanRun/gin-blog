package setting

import "time"

type setting struct {
	APP      app
	Server   server
	Database database
	Redis    redis
}

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
	DefaultPageSize uint     `json:"default_page_size"`
	MaxPageSize     uint     `json:"max_page_size"`
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
