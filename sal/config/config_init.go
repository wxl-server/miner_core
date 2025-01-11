package config

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/qcq1/common/json"
	"github.com/spf13/viper"
)

var Config = &LocalConfigStruct{}

type LocalConfigStruct struct {
	Nacos NacosConfig `yaml:"nacos"`
}

type NacosConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Namespace string `yaml:"namespace"`
}

func InitLocalConfig(ctx context.Context, configPath string) {
	vip := viper.New()
	vip.SetConfigFile(configPath)
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
	logger.CtxInfof(ctx, "[Init] init local config success, config = %v", json.MarshalWithoutError[string](Config))
}
