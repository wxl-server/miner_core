package main

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/server"
	"github.com/qcq1/common/env"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core/minercore"
	"miner_core/app"
	"miner_core/sal/config"
	"net"
)

type Handler struct {
	app *app.App
}

func NewHandler(app *app.App) *Handler {
	return &Handler{
		app: app,
	}
}

func (s *Handler) QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq) (r *miner_core.QueryJobListResp, err error) {
	return s.app.P.JobService.QueryJobList(ctx, req)
}

func runServer(ctx context.Context, config *config.AppConfig, handler *Handler) {
	options := make([]server.Option, 0)
	if env.IsBoe() {
		serverConfig := config.Server
		addr, err := net.ResolveTCPAddr(serverConfig.Network, serverConfig.HostPort)
		if err != nil {
			logger.CtxErrorf(ctx, "[Init] resolve tcp addr failed, err = %v", err)
			panic(err)
		}
		options = append(options, server.WithServiceAddr(addr))
	}
	svr := minercore.NewServer(handler, options...)
	if err := svr.Run(); err != nil {
		logger.CtxErrorf(ctx, "[Init] server run failed, err = %v", err)
		panic(err)
	}
}
