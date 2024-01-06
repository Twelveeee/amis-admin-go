package conf

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AppName string `yaml:"appName"`
	Env     string `yaml:"env"`

	HttpServer struct {
		Port         int `yaml:"port"`
		ReadTimeout  int `yaml:"readTimeout"`  // ms
		WriteTimeout int `yaml:"writeTimeout"` // ms
		IdleTimeout  int `yaml:"idleTimeout"`  // ms
	} `yaml:"httpServer"`
}

func ParseConfig(path string) (*Config, error) {
	var config Config
	if err := LoadYaml(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func LoadYaml(path string, v interface{}) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(file, v)
}

// func AppInit(Config)
