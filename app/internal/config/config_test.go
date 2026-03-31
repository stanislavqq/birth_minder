package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
	"github.com/gookit/goutil/testutil/assert"
	"reflect"
	"testing"
)

func TestGetConfigInstance(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "Success instance",
			want: &Config{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetConfigInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfigInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadConfigYML(t *testing.T) {
	var cfg Config

	config.WithOptions(config.ParseEnv, func(opt *config.Options) {
		// Указываем тег для сопоставления данных к полям структур. По умолчанию = mapstructure
		opt.DecoderConfig.TagName = yamlv3.Driver.Name()
	})
	config.AddDriver(yamlv3.Driver)

	err := config.LoadFiles("../../config.yml")
	assert.NoErr(t, err)

	err = config.BindStruct("", &cfg)
	assert.NoErr(t, err)

	assert.Eq(t, 1234567, cfg.TGBot.NotifyChat)
}
