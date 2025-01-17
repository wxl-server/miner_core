package service

import (
	"context"
	"miner_core/domain/converter"
	"miner_core/repo"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_core"
	"go.uber.org/dig"
)

type JobService interface {
	QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq) (r *miner_core.QueryJobListResp, err error)
}

type JobServiceImpl struct {
	p Param
}

type Param struct {
	dig.In
	repo.JobRepo
}

func NewJobService(p Param) JobService {
	return &JobServiceImpl{
		p: p,
	}
}

func (s *JobServiceImpl) QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq) (r *miner_core.QueryJobListResp, err error) {
	jobs, err := s.p.JobRepo.QueryJobList(ctx, converter.BuildQueryJobListReq(req))
	if err != nil {
		logger.CtxErrorf(ctx, "s.p.JobRepo.QueryJobList failed, err = %v", err)
		return nil, err
	}
	return &miner_core.QueryJobListResp{
		JobList: converter.JobDOs2DTOs(jobs),
	}, nil
}
