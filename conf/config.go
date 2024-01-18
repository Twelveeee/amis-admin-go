package conf

import (
	"fmt"
	"sync"

	"github.com/twelveeee/amis-admin-go/util"
)

var AppConf *Config

var confOnce sync.Once

type Config struct {
	AppName string `yaml:"appName"`
	Env     string `yaml:"env"`

	HttpServer struct {
		Port         int `yaml:"port"`
		ReadTimeout  int `yaml:"readTimeout"`  // ms
		WriteTimeout int `yaml:"writeTimeout"` // ms
		IdleTimeout  int `yaml:"idleTimeout"`  // ms
	} `yaml:"httpServer"`

	Mysql struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbName"`
	} `yaml:"mysql"`
}

func parseConfig(path string) (*Config, error) {
	var config Config
	if err := util.LoadYaml(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func InitConfig(path string) error {
	confOnce.Do(func() {
		conf, err := parseConfig(path)
		if err != nil {
			fmt.Println(err)
		}
		AppConf = conf
	})
	return nil
}

func GetAppConf() *Config {
	return AppConf
}
