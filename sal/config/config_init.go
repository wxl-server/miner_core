package config

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/luci/go-render/render"
	"github.com/qcq1/common/env"
	"github.com/spf13/viper"
)

var Config = &LocalConfigStruct{}

type LocalConfigStruct struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Name     string `yaml:"name"`
	HostPort string `yaml:"host_port"`
	Network  string `yaml:"network"`
}

func InitLocalConfig(ctx context.Context, configPath string, envStr env.Env) {
	vip := viper.New()
	vip.SetConfigFile(configPath + envStr + ".yaml")
	err := vip.ReadInConfig()
	if err != nil {
		logger.CtxErrorf(ctx, "[Init] init local config failed, err = %v", err)
		panic(err)
	}
	err = vip.Unmarshal(Config)
	if err != nil {
		logger.CtxErrorf(ctx, "[Init] unmarshal config failed, err = %v", err)
		panic(err)
	}
	logger.CtxInfof(ctx, "[Init] init local config success, config = %v", render.Render(Config))
}
