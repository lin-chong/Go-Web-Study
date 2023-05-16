package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var Config *AppConfig

type AppConfig struct {
	Http struct {
		Service string `yaml:"service"`
		Port    string `yaml:"port"`
	} `yaml:"http"`

	DataSource struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Address  string `yaml:"address"`
		DBName   string `yaml:"db_name"`
	} `yaml:"datasource"`
}

func LoadConfig() (*AppConfig, error) {
	ymlFile, err := os.ReadFile("config/config.yml")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	appConfig := new(AppConfig)
	err = yaml.Unmarshal(ymlFile, appConfig)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return appConfig, nil
}
