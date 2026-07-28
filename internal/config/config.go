package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
)

var cfg *Config

type Database struct {
	Migrations string   `yaml:"migrations"`
	Driver     string   `yaml:"driver"`
	Mysql      Mysql    `yaml:"mysql"`
	Sqlite     Sqlite   `yaml:"sqlite"`
	Postgres   Postgres `yaml:"postgres"`
}

type Postgres struct {
	Enabled  bool   `yaml:"enabled"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SslMode  string `yaml:"sslmode"`
}

type Mysql struct {
	Enabled  bool   `yaml:"enabled"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SslMode  string `yaml:"sslmode"`
}

type Sqlite struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
}

type TGBot struct {
	Token      string `yaml:"token"`
	NotifyChat int    `yaml:"notifyChat"`
}

type Config struct {
	Debug         bool     `yaml:"debug"`
	FormatMessage string   `yaml:"formatMessage"`
	Database      Database `yaml:"database"`
	CronRule      string   `yaml:"cronRule"`
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
