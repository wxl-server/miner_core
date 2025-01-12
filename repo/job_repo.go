package repo

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/dig"
	"miner_core/domain"
	"miner_core/domain/converter"
	"miner_core/sal/dao"
)

type JobRepo interface {
	QueryJobList(ctx context.Context, req *domain.QueryJobListReqDO) ([]domain.JobDO, error)
}

type JobRepoImpl struct {
	p Param
}
type Param struct {
	dig.In
	JobDal dao.JobDal
}

func NewJobRepo(p Param) JobRepo {
	return &JobRepoImpl{
		p: p,
	}
}

func (i JobRepoImpl) QueryJobList(ctx context.Context, req *domain.QueryJobListReqDO) ([]domain.JobDO, error) {
	jobList, err := i.p.JobDal.QueryJobList(ctx, converter.BuildJobWhereOpt(req))
	if err != nil {
		logger.CtxErrorf(ctx, "i.p.JobDal.QueryJobList failed, err = %v", err)
		return nil, err
	}
	return converter.JobPOs2DOs(jobList), nil
}
