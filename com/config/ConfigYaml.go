package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type Config struct {
	Name string
}

// 监控配置文件变化并热加载程序
func (c *Config) WatchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}

func (c *Config) Init() error {
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件,指定文件路径为相对路径./resource/dev/config.yaml
		viper.SetConfigFile(c.Name)
	} else {
		// 如果没有指定配置文件，则根据当前运行环境解析默认的配置文件
		viper.AddConfigPath(GetEnvConfig("dev"))
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")     // 设置配置文件格式为YAML
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix("HuJianJun") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}
	return nil
}

func InitConfig(cfg string) error {
 	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	if err := c.Init(); err != nil {
		return err
	}
	// 监控配置文件变化并热加载程序
	c.WatchConfig()
	log.Println("成功初始化文件")
	return nil
}

/**
根据运行环境读取相应的配置
安装golang环境设置运行环境变量GO_ENV=dev/st/ga
*/
func GetEnvConfig(param string) string {
	var ProjectConfigPath string
	var env string
	if(param==""){
		//获取当前环境
		env=os.Getenv("GO_ENV")
	}else{
		env=param
	}

	//var envStruct Env
	if env == "dev" {
		ProjectConfigPath = "resource/dev/"
	} else if env == "st" {
		ProjectConfigPath = "../resource/st/"
	} else {
		ProjectConfigPath = "../resource/ga/"
	}
	return ProjectConfigPath
}
