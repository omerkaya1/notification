package config

import (
	"github.com/omerkaya1/notification/internal/errors"
	"github.com/spf13/viper"
	"path"
)

type Config struct {
	Host      string `json:"host",yaml:"host",toml:"host"`
	Port      string `json:"port",yaml:"port",toml:"port"`
	User      string `json:"user",yaml:"user",toml:"user"`
	Password  string `json:"password",yaml:"password",toml:"password"`
	QueueName string `json:"queuename" yaml:"queuename" toml:"queuename"`
}

func InitConfig(cfgPath string) (*Config, error) {
	viper.SetConfigFile(cfgPath)

	cfgFileExt := path.Ext(cfgPath)
	if cfgFileExt == "" {
		return nil, errors.ErrBadConfigFile
	}
	viper.SetConfigType(cfgFileExt[1:])

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
