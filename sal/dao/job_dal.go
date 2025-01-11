package dao

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"miner_core/sal/dao/model"
	"miner_core/sal/dao/query"
)

type JobDal interface {
	QueryJobList(ctx context.Context, condition query.IJobPODo) ([]*model.JobPO, error)
}

type JobDalImpl struct {
}

func NewJobDal() JobDal {
	return &JobDalImpl{}
}

func (j JobDalImpl) QueryJobList(ctx context.Context, condition query.IJobPODo) ([]*model.JobPO, error) {
	jobPOS, err := condition.Find()
	if err != nil {
		logger.CtxErrorf(ctx, "condition.Find failed, err = %v", err)
		return nil, err
	}
	return jobPOS, nil
}
