package query

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/qcq1/common/gptr"
	"miner_core/common/constants"
	"miner_core/domain"
)

func BuildQueryJobListCondition(ctx context.Context, req *domain.QueryJobListReq) (condition IJobPODo) {
	jobPO := Q.JobPO
	condition = jobPO.Limit(int(req.PageSize))
	condition = condition.Offset(int(req.PageSize * req.PageNum))
	if req.OrderBy != nil {
		if field, ok := jobPO.GetFieldByName(gptr.Indirect(req.OrderBy)); ok {
			if req.Order != nil && gptr.Indirect(req.Order) == constants.Desc {
				condition = condition.Order(field.Desc())
			} else {
				condition = condition.Order(field)
			}
		} else {
			logger.CtxErrorf(ctx, "field not found, field = %v", field)
			if req.Order != nil && gptr.Indirect(req.Order) == constants.Desc {
				condition = condition.Order(jobPO.ID.Desc())
			} else {
				condition = condition.Order(jobPO.ID)
			}
		}
	}
	if req.ID != nil {
		condition = condition.Where(jobPO.ID.Eq(gptr.Indirect(req.ID)))
	}
	if req.CreatedBy != nil {
		condition = condition.Where(jobPO.CreatedBy.Eq(gptr.Indirect(req.CreatedBy)))
	}
	if req.CreatedAtStart != nil {
		condition = condition.Where(jobPO.CreatedAt.Gte(gptr.Indirect(req.CreatedAtStart)))
	}
	if req.CreatedAtEnd != nil {
		condition = condition.Where(jobPO.CreatedAt.Lte(gptr.Indirect(req.CreatedAtEnd)))
	}
	return condition
}
