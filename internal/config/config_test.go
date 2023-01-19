package config

import (
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
	type args struct {
		configYML string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadConfigYML(tt.args.configYML); (err != nil) != tt.wantErr {
				t.Errorf("ReadConfigYML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
