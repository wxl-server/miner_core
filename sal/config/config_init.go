package config

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/luci/go-render/render"
	"github.com/qcq1/common/env"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Server ServerConfig `yaml:"server"`
	Nacos  NacosConfig  `yaml:"nacos"`
	Mysql  MysqlConfig  `yaml:"mysql"`
}

type ServerConfig struct {
	Name     string `yaml:"name"`
	HostPort string `yaml:"host_port"`
	Network  string `yaml:"network"`
}

type NacosConfig struct {
	Host      string `yaml:"host"`
	Port      uint64 `yaml:"port"`
	Namespace string `yaml:"namespace"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

type MysqlConfig struct {
	Dsn string `yaml:"dsn"`
}

func InitAppConfig(ctx context.Context) *AppConfig {
	vip := viper.New()
	vip.SetConfigFile("conf/" + env.GetEnv() + ".yaml")
	err := vip.ReadInConfig()
	if err != nil {
		logger.CtxErrorf(ctx, "[Init] init local config failed, err = %v", err)
		panic(err)
	}
	Config := &AppConfig{}
	err = vip.Unmarshal(Config)
	if err != nil {
		logger.CtxErrorf(ctx, "[Init] unmarshal config failed, err = %v", err)
		panic(err)
	}
	logger.CtxInfof(ctx, "[Init] init local config success, config = %v", render.Render(Config))
	return Config
}
