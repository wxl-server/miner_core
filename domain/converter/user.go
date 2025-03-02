package converter

import (
	"miner_core/domain"
	"miner_core/sal/dao/generate/model"

	"github.com/wxl-server/common/gptr"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_core"
)

func SignUpReqDTO2DO(req *miner_core.SignUpReq) *domain.UserDO {
	return &domain.UserDO{
		Email:    req.Email,
		Password: req.Password,
	}
}

func UserDO2PO(do *domain.UserDO) *model.UserPO {
	return &model.UserPO{
		ID:        do.ID,
		Email:     do.Email,
		Password:  do.Password,
		Extra:     do.Extra,
		CreatedAt: do.CreatedAt,
		UpdatedAt: do.UpdatedAt,
	}
}

func UserPOs2DOs(pos []*model.UserPO) []domain.UserDO {
	dos := make([]domain.UserDO, 0, len(pos))
	for _, po := range pos {
		dos = append(dos, gptr.Indirect(UserPO2DO(po)))
	}
	return dos
}

func UserPO2DO(po *model.UserPO) *domain.UserDO {
	return &domain.UserDO{
		ID:        po.ID,
		Email:     po.Email,
		Password:  po.Password,
		Extra:     po.Extra,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
	}
}
