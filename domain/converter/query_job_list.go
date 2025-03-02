package converter

import (
	"miner_core/domain"
	"miner_core/sal/dao/generate/model"
	"strings"
	"time"

	"github.com/wxl-server/common/choose"
	"github.com/wxl-server/common/gptr"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_core"
)

func BuildQueryJobListReq(dto *miner_core.QueryJobListReq) *domain.QueryJobListReqDO {
	do := &domain.QueryJobListReqDO{
		PageNum:        dto.PageNum,
		PageSize:       dto.PageSize,
		ID:             dto.Id,
		Name:           dto.Name,
		CreatedBy:      dto.CreatedBy,
		CreatedAtStart: choose.If(dto.CreatedAtStart == nil, nil, gptr.Of(time.Unix(gptr.Indirect(dto.CreatedAtStart), 0))),
		CreatedAtEnd:   choose.If(dto.CreatedAtEnd == nil, nil, gptr.Of(time.Unix(gptr.Indirect(dto.CreatedAtEnd), 0))),
	}
	if dto.OrderBy != nil && gptr.Indirect(dto.OrderBy) != -1 {
		do.OrderBy = gptr.Of(strings.ToLower(dto.OrderBy.String()))
	}
	if dto.Order != nil && gptr.Indirect(dto.Order) != -1 {
		do.Order = gptr.Of(strings.ToLower(dto.Order.String()))
	}
	return do
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

func JobDOs2DTOs(dos []domain.JobDO, userIDs2DO map[int64]domain.UserDO) (dtos []*miner_core.Job) {
	dtos = make([]*miner_core.Job, 0, len(dos))
	for do := range dos {
		dtos = append(dtos, JobDO2DTO(dos[do], userIDs2DO[dos[do].CreatedBy], userIDs2DO[dos[do].UpdatedBy]))
	}
	return dtos
}

func JobDO2DTO(do domain.JobDO, creator domain.UserDO, updater domain.UserDO) *miner_core.Job {
	return &miner_core.Job{
		Id:          do.ID,
		Name:        do.Name,
		Description: do.Description,
		CreatedBy: &miner_core.User{
			Id:    do.CreatedBy,
			Email: creator.Email,
		},
		UpdatedBy: &miner_core.User{
			Id:    do.UpdatedBy,
			Email: updater.Email,
		},
		Extra:     do.Extra,
		CreatedAt: do.CreatedAt.Unix(),
		UpdatedAt: do.UpdatedAt.Unix(),
	}
}
