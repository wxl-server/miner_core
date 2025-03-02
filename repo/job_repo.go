package repo

import (
	"context"
	"miner_core/common/consts"
	"miner_core/domain"
	"miner_core/domain/converter"
	"miner_core/sal/dao/generate/query"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/common/gptr"
	"github.com/wxl-server/common/id_gen"
	"go.uber.org/dig"
)

type JobRepo interface {
	QueryJobList(ctx context.Context, req *domain.QueryJobListReqDO) ([]domain.JobDO, error)
	CreateJob(ctx context.Context, do *domain.JobDO) (int64, error)
	DeleteJob(ctx context.Context, id int64) error
}

type JobRepoImpl struct {
	p Param
}

type Param struct {
	dig.In
}

func NewJobRepo(p Param) JobRepo {
	return &JobRepoImpl{
		p: p,
	}
}

func (i JobRepoImpl) QueryJobList(ctx context.Context, req *domain.QueryJobListReqDO) ([]domain.JobDO, error) {
	req.PageNum -= 1
	po := query.Q.JobPO
	condition := po.WithContext(ctx).Limit(int(req.PageSize))
	condition = condition.Offset(int(req.PageSize * req.PageNum))
	if req.OrderBy != nil {
		if field, ok := po.GetFieldByName(gptr.Indirect(req.OrderBy)); ok {
			if req.Order != nil && gptr.Indirect(req.Order) == consts.Desc {
				condition = condition.Order(field.Desc())
			} else {
				condition = condition.Order(field)
			}
		} else {
			logger.CtxErrorf(ctx, "field not found, field = %v", field)
			if req.Order != nil && gptr.Indirect(req.Order) == consts.Desc {
				condition = condition.Order(po.ID.Desc())
			} else {
				condition = condition.Order(po.ID)
			}
		}
	} else {
		condition = condition.Order(po.UpdatedAt.Desc())
	}
	if req.ID != nil {
		condition = condition.Where(po.ID.Eq(gptr.Indirect(req.ID)))
	}
	if req.Name != nil {
		condition = condition.Where(po.Name.Like("%" + gptr.Indirect(req.Name) + "%"))
	}
	if req.CreatedBy != nil {
		condition = condition.Where(po.CreatedBy.Eq(gptr.Indirect(req.CreatedBy)))
	}
	if req.CreatedAtStart != nil {
		condition = condition.Where(po.CreatedAt.Gte(gptr.Indirect(req.CreatedAtStart)))
	}
	if req.CreatedAtEnd != nil {
		condition = condition.Where(po.CreatedAt.Lte(gptr.Indirect(req.CreatedAtEnd)))
	}

	jobList, err := condition.Find()
	if err != nil {
		logger.CtxErrorf(ctx, "condition.Find failed, err = %v", err)
		return nil, err
	}
	return converter.JobPOs2DOs(jobList), nil
}

func (i JobRepoImpl) CreateJob(ctx context.Context, do *domain.JobDO) (id int64, err error) {
	po := query.Q.JobPO
	do.ID, err = id_gen.NextID()
	err = po.WithContext(ctx).Create(converter.JobDO2PO(do))
	if err != nil {
		logger.CtxErrorf(ctx, "po.WithContext(ctx).Create failed, err = %v", err)
		return 0, err
	}
	return do.ID, nil
}

func (i JobRepoImpl) DeleteJob(ctx context.Context, id int64) error {
	po := query.Q.JobPO
	_, err := po.WithContext(ctx).Where(po.ID.Eq(id)).Delete()
	if err != nil {
		logger.CtxErrorf(ctx, "po.WithContext(ctx).Where(po.ID.Eq(id)).Delete failed, err = %v", err)
		return err
	}
	return nil
}
