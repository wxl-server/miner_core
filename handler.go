package main

import (
	"context"
	"miner_core/service"

	"github.com/wxl-server/idl_gen/kitex_gen/miner_core"
	"go.uber.org/dig"
)

var handler miner_core.MinerCore

type Handler struct {
	p Param
}

type Param struct {
	dig.In
	JobService       service.JobService
	UserService      service.UserService
	IndicatorService service.IndicatorService
	TaskService      service.TaskService
}

func NewHandler(p Param) {
	handler = &Handler{
		p: p,
	}
}

func (s *Handler) SignUp(ctx context.Context, req *miner_core.SignUpReq) (r *miner_core.SignUpResp, err error) {
	return s.p.UserService.SignUp(ctx, req)
}

func (s *Handler) Login(ctx context.Context, req *miner_core.LoginReq) (r *miner_core.LoginResp, err error) {
	return s.p.UserService.Login(ctx, req)
}

func (s *Handler) QueryUserList(ctx context.Context, req *miner_core.QueryUserListReq) (r *miner_core.QueryUserListResp, err error) {
	return s.p.UserService.QueryUserList(ctx, req)
}

func (s *Handler) QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq) (r *miner_core.QueryJobListResp, err error) {
	return s.p.JobService.QueryJobList(ctx, req)
}

func (s *Handler) CreateJob(ctx context.Context, req *miner_core.CreateJobReq) (r *miner_core.CreateJobResp, err error) {
	return s.p.JobService.CreateJob(ctx, req)
}

func (s *Handler) DeleteJob(ctx context.Context, req *miner_core.DeleteJobReq) (r *miner_core.DeleteJobResp, err error) {
	return s.p.JobService.DeleteJob(ctx, req)
}

func (s *Handler) QueryIndicatorList(ctx context.Context, req *miner_core.QueryIndicatorListReq) (r *miner_core.QueryIndicatorListResp, err error) {
	return s.p.IndicatorService.QueryIndicatorList(ctx, req)
}

func (s *Handler) RunTask(ctx context.Context, req *miner_core.RunTaskReq) (r *miner_core.RunTaskResp, err error) {
	return s.p.TaskService.RunTask(ctx, req)
}
