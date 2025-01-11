package main

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/qcq1/common/env"
	miner_core "github.com/qcq1/rpc_miner_core/kitex_gen/miner_core/itemservice"
	"miner_core/sal/config"
)

func main() {
	ctx := context.Background()
	Init(ctx)

	svr := miner_core.NewServer(
		new(ItemServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "miner_core"}),
		server.WithRegistry(registry.NewNacosRegistry(InitNacosClient(ctx))),
	)
	if err := svr.Run(); err != nil {
		logger.CtxErrorf(ctx, "[Init] server run failed, err = %v", err)
		panic(err)
	}
}

func Init(ctx context.Context) {
	config.InitLocalConfig(ctx, "conf/", env.GetEnv())
}

func InitNacosClient(ctx context.Context) naming_client.INamingClient {
	nacosConfig := config.Config.Nacos
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(nacosConfig.Host, nacosConfig.Port, constant.WithGrpcPort(nacosConfig.GrpcPort)),
	}

	cc := constant.ClientConfig{
		NamespaceId: nacosConfig.Namespace,
		Username:    nacosConfig.Username,
		Password:    nacosConfig.Password,
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		logger.CtxErrorf(ctx, "[Init] nacos init failed, err = %v", err)
		panic(err)
	}
	return cli
}
