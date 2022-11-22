package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var Conf = new(Config)

func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("读取应用目录失败:%s", err))
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/")
	// 读取配置信息
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败:%s", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 将读取的配置信息保存至全局变量Conf
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("初始化配置文件失败:%s", err))
		}

	})
	// 将读取的配置信息保存至全局变量Conf
	if err = viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("初始化配置文件失败:%s", err))
	}
}

type Config struct {
	System    *System    `mapstructure:"system" json:"system"`
	Log       *Log       `mapstructure:"log" json:"log"`
	Redis     *Redis     `mapstructure:"redis" json:"redis"`
	MySql     *MySql     `mapstructure:"mysql" json:"mysql"`
	Casbin    *Casbin    `mapstructure:"casbin" json:"casbin"`
	Jwt       *Jwt       `mapstructure:"jwt" json:"jwt"`
	RateLimit *RateLimit `mapstructure:"rate-limit" json:"rateLimit"`
	Email     *Email     `mapstructure:"email" json:"email"`
}
type System struct {
	Mode string `mapstructure:"mode" json:"mode"`
	Port int    `mapstructure:"port" json:"port"`
}
type Log struct {
	Level      string `mapstructure:"level" json:"level"`
	Path       string `mapstructure:"path" json:"path"`
	MaxSize    int    `mapstructure:"max-size" json:"maxSize"`
	MaxBackups int    `mapstructure:"max-backups" json:"maxBackups"`
	MaxAge     int    `mapstructure:"max-age" json:"maxAge"`
	Compress   bool   `mapstructure:"compress" json:"compress"`
}
type MySql struct {
	UserName    string `mapstructure:"username" json:"username"`
	Password    string `mapstructure:"password" json:"password"`
	Database    string `mapstructure:"database" json:"database"`
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Query       string `mapstructure:"query" json:"query"`
	LogMode     bool   `mapstructure:"log-mode" json:"logMode"`
	TablePrefix string `mapstructure:"table-prefix" json:"tablePrefix"`
	Charset     string `mapstructure:"charset" json:"charset"`
	Collation   string `mapstructure:"collation" json:"collation"`
}
type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath"`
}
type Jwt struct {
	Realm      string `mapstructure:"realm" json:"realm"`
	Key        string `mapstructure:"key" json:"key"`
	TimeOut    int    `mapstructure:"timeout" json:"timeout"`
	MaxRefresh int    `mapstructure:"max-refresh" json:"maxRefresh"`
}
type RateLimit struct {
	FillInterval int `mapstructure:"fill-interval" json:"fillInterval"`
	Capacity     int `mapstructure:"capacity" json:"capacity"`
}
type Email struct {
	Port int    `mapstructure:"port" json:"port"`
	User string `mapstructure:"user" json:"user"`
	From string `mapstructure:"from" json:"from"`
	Host string `mapstructure:"host" json:"host"`
	Pass string `mapstructure:"pass" json:"pass"`
}
type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr"`
	Password string `mapstructure:"password" json:"password"`
	DB       int    `mapstructure:"db" json:"db"`
}
