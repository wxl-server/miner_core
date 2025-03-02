package repo

import (
	"context"
	"miner_core/domain"
	"miner_core/domain/converter"
	"miner_core/sal/dao/generate/query"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/common/gslice"
	"github.com/wxl-server/common/id_gen"

	"go.uber.org/dig"
)

type UserRepo interface {
	CountUser(ctx context.Context, email string) (count int64, err error)
	CreateUser(ctx context.Context, do *domain.UserDO) (id int64, err error)
	QueryUser(ctx context.Context, where domain.QueryUserReqDO) (dos map[int64]domain.UserDO, err error)
}

type param struct {
	dig.In
}

type UserRepoImpl struct {
	p param
}

func NewUserRepo(p param) UserRepo {
	return &UserRepoImpl{
		p: p,
	}
}

func (u UserRepoImpl) QueryUser(ctx context.Context, where domain.QueryUserReqDO) (dos map[int64]domain.UserDO, err error) {
	po := query.Q.UserPO
	condition := po.WithContext(ctx)
	if where.ID != nil {
		condition = condition.Where(po.ID.Eq(*where.ID))
	} else if len(where.IDs) > 0 {
		condition = condition.Where(po.ID.In(where.IDs...))
	}
	if where.Email != nil {
		condition = condition.Where(po.Email.Eq(*where.Email))
	} else if len(where.Emails) > 0 {
		condition = condition.Where(po.Email.In(where.Emails...))
	}
	userPOs, err := condition.Find()
	if err != nil {
		logger.CtxErrorf(ctx, "QueryUser failed, err = %v", err)
		return nil, err
	}

	users := converter.UserPOs2DOs(userPOs)
	userIDs2DO := gslice.ToMap(users, func(user domain.UserDO) (int64, domain.UserDO) {
		return user.ID, user
	})
	return userIDs2DO, nil
}

func (u UserRepoImpl) CountUser(ctx context.Context, email string) (count int64, err error) {
	po := query.Q.UserPO
	count, err = po.WithContext(ctx).Where(po.Email.Eq(email)).Count()
	if err != nil {
		logger.CtxErrorf(ctx, "CountUser failed, err = %v", err)
		return 0, err
	}
	return count, nil
}

func (u UserRepoImpl) CreateUser(ctx context.Context, do *domain.UserDO) (id int64, err error) {
	po := query.Q.UserPO
	do.ID, err = id_gen.NextID()
	if err != nil {
		logger.CtxErrorf(ctx, "generate id failed, err = %v", err)
		return 0, err
	}
	err = po.WithContext(ctx).Create(converter.UserDO2PO(do))
	if err != nil {
		logger.CtxErrorf(ctx, "CreateUser failed, err = %v", err)
		return 0, err
	}
	return do.ID, nil
}
