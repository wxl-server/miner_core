package where

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/common/gptr"
	"miner_core/common/constants"
	query2 "miner_core/sal/dao/generator/query"
	"time"
)

type JobWhereOpt struct {
	PageNum  int64
	PageSize int64
	OrderBy  *string
	Order    *constants.Order

	ID             *int64
	CreatedBy      *int64
	CreatedAtStart *time.Time
	CreatedAtEnd   *time.Time
}

func JobWhereOpt2Condition(ctx context.Context, whereOpt *JobWhereOpt) (condition query2.IJobPODo) {
	po := query2.Q.JobPO
	condition = po.WithContext(ctx).Limit(int(whereOpt.PageSize))
	condition = condition.Offset(int(whereOpt.PageSize * whereOpt.PageNum))
	if whereOpt.OrderBy != nil {
		if field, ok := po.GetFieldByName(gptr.Indirect(whereOpt.OrderBy)); ok {
			if whereOpt.Order != nil && gptr.Indirect(whereOpt.Order) == constants.Desc {
				condition = condition.Order(field.Desc())
			} else {
				condition = condition.Order(field)
			}
		} else {
			logger.CtxErrorf(ctx, "field not found, field = %v", field)
			if whereOpt.Order != nil && gptr.Indirect(whereOpt.Order) == constants.Desc {
				condition = condition.Order(po.ID.Desc())
			} else {
				condition = condition.Order(po.ID)
			}
		}
	}
	if whereOpt.ID != nil {
		condition = condition.Where(po.ID.Eq(gptr.Indirect(whereOpt.ID)))
	}
	if whereOpt.CreatedBy != nil {
		condition = condition.Where(po.CreatedBy.Eq(gptr.Indirect(whereOpt.CreatedBy)))
	}
	if whereOpt.CreatedAtStart != nil {
		condition = condition.Where(po.CreatedAt.Gte(gptr.Indirect(whereOpt.CreatedAtStart)))
	}
	if whereOpt.CreatedAtEnd != nil {
		condition = condition.Where(po.CreatedAt.Lte(gptr.Indirect(whereOpt.CreatedAtEnd)))
	}
	return condition
}
