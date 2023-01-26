package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
)

var cfg *Config

type Database struct {
	Enabled    bool   `yaml:"enabled"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Migrations string `yaml:"migrations"`
	Name       string `yaml:"name"`
	SslMode    string `yaml:"sslmode"`
	Driver     string `yaml:"driver"`
}

type TGBot struct {
	Token      string `yaml:"token"`
	NotifyChat int    `yaml:"notifyChat"`
}

type Config struct {
	Debug         bool     `yaml:"debug"`
	FormatMessage string   `yaml:"formatMessage"`
	Database      Database `yaml:"database"`
	TGBot         TGBot    `yaml:"tgbot"`
}

func GetConfigInstance() *Config {
	if cfg != nil {
		return cfg
	}

	return &Config{}
}

func ReadConfigYML(configYML string) error {
	if cfg != nil {
		return nil
	}

	config.WithOptions(config.ParseEnv)
	config.WithTagName(yamlv3.Driver.Name())
	config.AddDriver(yamlv3.Driver)

	err := config.LoadFiles(configYML)
	if err != nil {
		return err
	}

	// привяжем структуру без ключа - так как у нас его нет
	if err := config.BindStruct("", &cfg); err != nil {
		return err
	}

	return nil
}
