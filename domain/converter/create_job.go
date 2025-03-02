package converter

import (
	"miner_core/domain"
	"miner_core/sal/dao/generate/model"
)

func JobDO2PO(do *domain.JobDO) *model.JobPO {
	return &model.JobPO{
		ID:          do.ID,
		Name:        do.Name,
		Description: do.Description,
		CreatedBy:   do.CreatedBy,
		UpdatedBy:   do.UpdatedBy,
		Extra:       do.Extra,
		CreatedAt:   do.CreatedAt,
		UpdatedAt:   do.UpdatedAt,
	}
}
