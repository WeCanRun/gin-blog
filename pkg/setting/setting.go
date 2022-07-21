package setting

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	Setting  = setting{}
	Redis    = redis{}
	Database = database{}
	Server   = server{}
	APP      = app{}

	DefaultConfigFile     = "/conf/app-%s.yaml"
	DefaultConfigFileType = "yaml"
	Env                   = flag.String("env", "test", "指定环境")
	Configuration         = flag.String("configuration", "", "指定配置文件")
)

func Setup(path string) {
	_type := DefaultConfigFileType

	if path == "" && *Configuration != "" {
		path = *Configuration
	}

	if path == "" {
		pwd, _ := os.Getwd()
		base := filepath.Dir(filepath.Dir(pwd))
		file := fmt.Sprintf(DefaultConfigFile, *Env)

		path = base + file
		if runtime.GOOS == "windows" {
			path = strings.ReplaceAll(path, "/", "\\")
		}

	}

	if path != "" {
		split := strings.Split(path, ".")
		_type = split[len(split)-1]
	}

	c := Config{
		Name: path,
		Type: _type,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		log.Fatalf("初始化配置文件失败，%v", err)
	}
	// 监控配置文件变化并热加载程序
	c.watchConfig()
}

type Config struct {
	Name string
	Type string
}

func (c *Config) initConfig() error {
	log.Printf("配置文件路径: %s\n", c.Name)
	viper.SetConfigFile(c.Name)
	viper.SetConfigType(c.Type)
	// 读取匹配的环境变量
	viper.AutomaticEnv()
	// 读取环境变量的前缀为 BLOG
	viper.SetEnvPrefix("BLOG")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fail to parse '%s': %v", c.Name, err)
		return err
	}

	c.loadData()

	log.Printf("Setting: %#v", Setting)

	return nil
}

func (c *Config) loadData() {
	if err := viper.Unmarshal(&Setting); err != nil {
		log.Fatalf("laod setting fail:  %v", err)
	}

	Database = Setting.Database
	Redis = Setting.Redis
	Server = Setting.Server
	APP = Setting.APP

}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config file changed: %s\n", e.Name)
		if err := viper.Unmarshal(&Setting); err != nil {
			log.Printf("err: %v\n", err)
		}

		Redis = Setting.Redis
		Database = Setting.Database
		Server = Setting.Server
		APP = Setting.APP
	})
}

func loadRedis() {
	if err := viper.Unmarshal(&Redis); err != nil {
		log.Fatalf("load Redis fail:  %v", err)
	}

	Redis.IdleTimeout *= time.Second
}

func loadDataBase() {
	if err := viper.Unmarshal(&Database); err != nil {
		log.Fatalf("laod Database fail:  %v", err)
	}
}

func loadServer() {
	if err := viper.Unmarshal(&Server); err != nil {
		log.Fatalf("load Server fail:  %v", err)
	}
	Server.ReadTimeout *= time.Second
	Server.WriteTimeout *= time.Second
}

func loadApp() {
	if err := viper.Unmarshal(&APP); err != nil {
		log.Fatalf("load APP fail: %v", err)
	}

	APP.ImageMaxSize *= 1024 * 1024
}
