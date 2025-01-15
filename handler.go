package main

import (
	"context"
	"miner_core/app"

	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"
)

var handler miner_core.MinerCore

type Handler struct {
	app *app.App
}

func NewHandler(app *app.App) {
	handler = &Handler{
		app: app,
	}
}

func (s *Handler) QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq) (r *miner_core.QueryJobListResp, err error) {
	return s.app.P.JobService.QueryJobList(ctx, req)
}
