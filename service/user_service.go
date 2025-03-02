package service

import (
	"context"
	"miner_core/biz_error"
	"miner_core/domain"
	"miner_core/domain/converter"
	"miner_core/repo"
	"miner_core/sal/jwt"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/common/gmap"
	"github.com/wxl-server/common/gptr"
	"github.com/wxl-server/common/gslice"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_core"
	"go.uber.org/dig"
)

type UserService interface {
	SignUp(ctx context.Context, req *miner_core.SignUpReq) (resp *miner_core.SignUpResp, err error)
	Login(ctx context.Context, req *miner_core.LoginReq) (resp *miner_core.LoginResp, err error)
	QueryUserList(ctx context.Context, req *miner_core.QueryUserListReq) (*miner_core.QueryUserListResp, error)
}

type UserServiceParam struct {
	dig.In
	UserRepo repo.UserRepo
}

type UserServiceImpl struct {
	p UserServiceParam
}

func NewUserService(p UserServiceParam) UserService {
	return &UserServiceImpl{
		p: p,
	}
}

func (u UserServiceImpl) SignUp(ctx context.Context, req *miner_core.SignUpReq) (resp *miner_core.SignUpResp, err error) {
	count, err := u.p.UserRepo.CountUser(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		logger.CtxErrorf(ctx, "sign up failed, email has been used, email = %s", req.Email)
		return nil, biz_error.SignUpError
	}
	id, err := u.p.UserRepo.CreateUser(ctx, converter.SignUpReqDTO2DO(req))
	if err != nil {
		logger.CtxErrorf(ctx, "CreateUser failed, err = %v", err)
		return nil, err
	}
	return &miner_core.SignUpResp{
		Id: id,
	}, nil
}

func (u UserServiceImpl) Login(ctx context.Context, req *miner_core.LoginReq) (resp *miner_core.LoginResp, err error) {
	id2DOs, err := u.p.UserRepo.QueryUser(ctx, domain.QueryUserReqDO{
		Email: gptr.Of(req.Email),
	})
	if err != nil {
		logger.CtxErrorf(ctx, "QueryUser failed, err = %v", err)
		return nil, biz_error.LoginError
	}

	values := gmap.Values(id2DOs)
	if len(values) == 0 {
		logger.CtxErrorf(ctx, "login failed, email not found, email = %s", req.Email)
		return nil, biz_error.LoginError
	}
	do := values[0]
	if do.Password != req.Password {
		logger.CtxErrorf(ctx, "login failed, password is wrong")
		return nil, biz_error.LoginError2
	}
	token, err := jwt.GenerateToken(ctx, do.ID)
	if err != nil {
		logger.CtxErrorf(ctx, "login failed, err = %v", err)
		return nil, biz_error.LoginError
	}
	return &miner_core.LoginResp{
		Token: token,
	}, nil
}

func (u UserServiceImpl) QueryUserList(ctx context.Context, req *miner_core.QueryUserListReq) (*miner_core.QueryUserListResp, error) {
	users, err := u.p.UserRepo.QueryUser(ctx, domain.QueryUserReqDO{})
	if err != nil {
		logger.CtxErrorf(ctx, "QueryUser failed, err = %v", err)
		return nil, err
	}
	return &miner_core.QueryUserListResp{
		UserList: gslice.Map(gmap.Values(users), func(user domain.UserDO) *miner_core.User {
			return &miner_core.User{
				Id:    user.ID,
				Email: user.Email,
			}
		}),
	}, nil
}
