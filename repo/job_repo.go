package repo

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/common/gptr"
	"go.uber.org/dig"
	"miner_core/common/constants"
	"miner_core/domain"
	"miner_core/domain/converter"
	query2 "miner_core/sal/dao/generator/query"
)

type JobRepo interface {
	QueryJobList(ctx context.Context, reqDO *domain.QueryJobListReqDO) ([]domain.JobDO, error)
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

func (i JobRepoImpl) QueryJobList(ctx context.Context, reqDO *domain.QueryJobListReqDO) ([]domain.JobDO, error) {
	po := query2.Q.JobPO
	condition := po.WithContext(ctx).Limit(int(reqDO.PageSize))
	condition = condition.Offset(int(reqDO.PageSize * reqDO.PageNum))
	if reqDO.OrderBy != nil {
		if field, ok := po.GetFieldByName(gptr.Indirect(reqDO.OrderBy)); ok {
			if reqDO.Order != nil && gptr.Indirect(reqDO.Order) == constants.Desc {
				condition = condition.Order(field.Desc())
			} else {
				condition = condition.Order(field)
			}
		} else {
			logger.CtxErrorf(ctx, "field not found, field = %v", field)
			if reqDO.Order != nil && gptr.Indirect(reqDO.Order) == constants.Desc {
				condition = condition.Order(po.ID.Desc())
			} else {
				condition = condition.Order(po.ID)
			}
		}
	}
	if reqDO.ID != nil {
		condition = condition.Where(po.ID.Eq(gptr.Indirect(reqDO.ID)))
	}
	if reqDO.CreatedBy != nil {
		condition = condition.Where(po.CreatedBy.Eq(gptr.Indirect(reqDO.CreatedBy)))
	}
	if reqDO.CreatedAtStart != nil {
		condition = condition.Where(po.CreatedAt.Gte(gptr.Indirect(reqDO.CreatedAtStart)))
	}
	if reqDO.CreatedAtEnd != nil {
		condition = condition.Where(po.CreatedAt.Lte(gptr.Indirect(reqDO.CreatedAtEnd)))
	}
	jobPOS, err := condition.Find()
	if err != nil {
		logger.CtxErrorf(ctx, "condition.Find failed, err = %v", err)
		return nil, err
	}
	return converter.JobPOs2DOs(jobPOS), nil
}
