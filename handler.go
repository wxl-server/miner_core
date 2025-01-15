package main

import (
	"context"
	"miner_core/service"

	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"
	"go.uber.org/dig"
)

var handler miner_core.MinerCore

type Handler struct {
	p Param
}

type Param struct {
	dig.In
	JobService service.JobService
}

func NewHandler(p Param) {
	handler = &Handler{
		p: p,
	}
}

func (s *Handler) QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq) (r *miner_core.QueryJobListResp, err error) {
	return s.p.JobService.QueryJobList(ctx, req)
}
