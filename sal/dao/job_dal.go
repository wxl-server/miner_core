package dao

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"miner_core/sal/dao/generate/model"
	"miner_core/sal/dao/where"
)

type JobDal interface {
	QueryJobList(ctx context.Context, whereOpt *where.JobWhereOpt) ([]*model.JobPO, error)
}

type JobDalImpl struct {
}

func NewJobDal() JobDal {
	return &JobDalImpl{}
}

func (j JobDalImpl) QueryJobList(ctx context.Context, whereOpt *where.JobWhereOpt) ([]*model.JobPO, error) {
	jobPOS, err := where.JobWhereOpt2Condition(ctx, whereOpt).Find()
	if err != nil {
		logger.CtxErrorf(ctx, "condition.Find failed, err = %v", err)
		return nil, err
	}
	return jobPOS, nil
}
