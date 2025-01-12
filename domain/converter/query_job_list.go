package converter

import (
	"github.com/qcq1/common/choose"
	"github.com/qcq1/common/gptr"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"
	"miner_core/domain"
	"miner_core/sal/dao/generate/model"
	"miner_core/sal/dao/where"
	"strings"
	"time"
)

func BuildQueryJobListReq(dto *miner_core.QueryJobListReq) *domain.QueryJobListReqDO {
	do := &domain.QueryJobListReqDO{
		PageNum:        dto.PageNum,
		PageSize:       dto.PageSize,
		ID:             dto.Id,
		CreatedBy:      dto.CreatedBy,
		CreatedAtStart: choose.If(dto.CreatedAtStart == nil, nil, gptr.Of(time.Unix(gptr.Indirect(dto.CreatedAtStart), 0))),
		CreatedAtEnd:   choose.If(dto.CreatedAtEnd == nil, nil, gptr.Of(time.Unix(gptr.Indirect(dto.CreatedAtEnd), 0))),
	}
	if dto.OrderBy != nil {
		do.OrderBy = gptr.Of(strings.ToLower(dto.OrderBy.String()))
	}
	if dto.Order != nil {
		do.Order = gptr.Of(strings.ToLower(dto.Order.String()))
	}
	return do
}

func BuildJobWhereOpt(do *domain.QueryJobListReqDO) *where.JobWhereOpt {
	return &where.JobWhereOpt{
		PageNum:        do.PageNum,
		PageSize:       do.PageSize,
		OrderBy:        do.OrderBy,
		Order:          do.Order,
		ID:             do.ID,
		CreatedBy:      do.CreatedBy,
		CreatedAtStart: do.CreatedAtStart,
	}
}

func JobPOs2DOs(pos []*model.JobPO) (dos []domain.JobDO) {
	dos = make([]domain.JobDO, 0, len(pos))
	for po := range pos {
		dos = append(dos, JobPO2DO(pos[po]))
	}
	return dos
}

func JobPO2DO(po *model.JobPO) domain.JobDO {
	return domain.JobDO{
		ID:          po.ID,
		Name:        po.Name,
		Description: po.Description,
		CreatedBy:   po.CreatedBy,
		UpdatedBy:   po.UpdatedBy,
		Extra:       po.Extra,
		CreatedAt:   po.CreatedAt,
		UpdatedAt:   po.UpdatedAt,
	}
}

func JobDOs2DTOs(dos []domain.JobDO) (dtos []*miner_core.Job) {
	dtos = make([]*miner_core.Job, 0, len(dos))
	for do := range dos {
		dtos = append(dtos, JobDO2DTO(dos[do]))
	}
	return dtos
}

func JobDO2DTO(do domain.JobDO) *miner_core.Job {
	return &miner_core.Job{
		Id:          do.ID,
		Name:        do.Name,
		Description: do.Description,
		CreatedBy:   do.CreatedBy,
		UpdatedBy:   do.UpdatedBy,
		Extra:       do.Extra,
		CreatedAt:   do.CreatedAt.Unix(),
		UpdatedAt:   do.UpdatedAt.Unix(),
	}
}
