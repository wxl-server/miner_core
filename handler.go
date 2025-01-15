package main

import (
	"context"
	"miner_core/app"
	"miner_core/sal/config"
	"net"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/qcq1/common/env"
	"github.com/qcq1/common/render"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core/minercore"
)

type Handler struct {
	app *app.App
}

func NewHandler(app *app.App) miner_core.MinerCore {
	return &Handler{
		app: app,
	}
}

func (s *Handler) QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq) (r *miner_core.QueryJobListResp, err error) {
	return s.app.P.JobService.QueryJobList(ctx, req)
}

func runServer(ctx context.Context, config *config.AppConfig, handler miner_core.MinerCore) {
	serverConfig := config.Server
	nacosConfig := config.Nacos
	options := make([]server.Option, 0)

	// boe 环境下指定服务地址
	if env.IsBoe() {
		addr, err := net.ResolveTCPAddr(serverConfig.Network, serverConfig.HostPort)
		if err != nil {
			logger.CtxErrorf(ctx, "[Init] resolve tcp addr failed, err = %v", err)
			panic(err)
		}
		options = append(options, server.WithServiceAddr(addr))
	}

	// 注册服务到nacos
	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				NamespaceId: nacosConfig.Namespace,
				Username:    nacosConfig.Username,
				Password:    nacosConfig.Password,
			},
			ServerConfigs: []constant.ServerConfig{
				*constant.NewServerConfig(nacosConfig.Host, nacosConfig.Port),
			},
		},
	)
	if err != nil {
		logger.CtxErrorf(ctx, "[Init] new nacos client failed, err = %v", err)
		panic(err)
	}
	options = append(options, server.WithRegistry(registry.NewNacosRegistry(cli)))
	options = append(options, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serverConfig.Name}))

	// 日志中间件
	options = append(options, server.WithMiddleware(LogMiddleware))

	// 启动服务
	svr := minercore.NewServer(handler, options...)
	if err = svr.Run(); err != nil {
		logger.CtxErrorf(ctx, "[Init] server run failed, err = %v", err)
		panic(err)
	}
}

func LogMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		if arg, ok := request.(utils.KitexArgs); ok {
			if req := arg.GetFirstArgument(); req != nil {
				logger.CtxInfof(ctx, "Get request = %v", render.Render(req))
			}
		}
		err := next(ctx, request, response)
		if result, ok := response.(utils.KitexResult); ok {
			if resp := result.GetResult(); resp != nil {
				logger.CtxInfof(ctx, "Send response = %v", render.Render(resp))
			}
		}
		return err
	}
}
