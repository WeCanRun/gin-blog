package setting

import (
	"flag"
	"fmt"
	"github.com/WeCanRun/gin-blog/pkg/file"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"
)

var (
	_setting              = &Setting{}
	DefaultConfigFile     = "conf/app-%s.yaml"
	DefaultConfigFileType = "yaml"
	Env                   = flag.String("env", "test", "指定环境")
	Configuration         = flag.String("configuration", "", "指定配置文件")
)

func Setup(path string) *Setting {
	_type := DefaultConfigFileType

	if path == "" && *Configuration != "" {
		path = *Configuration
	}

	if path == "" {
		f := fmt.Sprintf(DefaultConfigFile, *Env)
		path = file.CoverToAbs(f)
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
	return _setting
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

	log.Printf("_setting: %#v", _setting)

	return nil
}

func (c *Config) loadData() {
	if err := viper.Unmarshal(_setting); err != nil {
		log.Fatalf("laod Setting fail:  %v", err)
	}

	_setting.Redis.IdleTimeout *= time.Second
	_setting.Server.ReadTimeout *= time.Second
	_setting.Server.WriteTimeout *= time.Second
	_setting.APP.ImageMaxSize *= 1024 * 1024

}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		c.loadData()
	})
}
