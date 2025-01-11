package config

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/qcq1/common/env"
	"github.com/qcq1/common/json"
	"github.com/spf13/viper"
)

var Config = &LocalConfigStruct{}

type LocalConfigStruct struct {
	Nacos NacosConfig `yaml:"nacos"`
}

type NacosConfig struct {
	Host      string `yaml:"host"`
	Port      uint64 `yaml:"port"`
	GrpcPort  uint64 `yaml:"grpc_port"`
	Namespace string `yaml:"namespace"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
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
	logger.CtxInfof(ctx, "[Init] init local config success, config = %v", json.MarshalWithoutError[string](Config))
}
