package main

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/v2/registry"
	"github.com/qcq1/common/env"
	miner_core "github.com/qcq1/rpc_miner_core/kitex_gen/miner_core/itemservice"
	"miner_core/sal/config"
	"miner_core/sal/nacos"
	"net"
)

func main() {
	ctx := context.Background()
	Init(ctx)
	RunServer(ctx)
}

func Init(ctx context.Context) {
	config.InitLocalConfig(ctx, "conf/", env.GetEnv())
}

func RunServer(ctx context.Context) {
	serverConfig := config.Config.Server
	options := make([]server.Option, 0)
	if env.IsBoe() {
		addr, err := net.ResolveTCPAddr(serverConfig.Network, serverConfig.HostPort)
		if err != nil {
			logger.CtxErrorf(ctx, "[Init] resolve tcp addr failed, err = %v", err)
			panic(err)
		}
		options = append(options, server.WithServiceAddr(addr))
	}
	options = append(options, []server.Option{
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serverConfig.Name}),
		server.WithRegistry(registry.NewNacosRegistry(nacos.Register2Nacos(ctx))),
	}...)
	svr := miner_core.NewServer(
		new(ItemServiceImpl),
		options...,
	)
	if err := svr.Run(); err != nil {
		logger.CtxErrorf(ctx, "[Init] server run failed, err = %v", err)
		panic(err)
	}
}
