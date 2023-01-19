package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
)

var cfg *Config

type Project struct {
	Debug bool `yaml:"debug"`
}

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
	Token string `yaml:"token"`
}

type Config struct {
	Database Database `yaml:"database"`
	Project  Project  `yaml:"project"`
	TGBot    TGBot    `yaml:"tgbot"`
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

	// установим опции что парсим переменные окружения и будет yaml
	config.WithOptions(config.ParseEnv, func(opt *config.Options) {
		// Указываем тег для сопоставления данных к полям структур. По умолчанию = mapstructure
		opt.DecoderConfig.TagName = yamlv3.Driver.Name()
	})
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
