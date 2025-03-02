package service

import (
	"context"
	"miner_core/biz_error"
	"miner_core/domain"
	"miner_core/domain/converter"
	"miner_core/repo"
	"miner_core/sal/jwt"
	"strconv"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/common/gptr"
	"github.com/wxl-server/common/gslice"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_core"
	"go.uber.org/dig"
)

type JobService interface {
	QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq) (r *miner_core.QueryJobListResp, err error)
	CreateJob(ctx context.Context, req *miner_core.CreateJobReq) (r *miner_core.CreateJobResp, err error)
	DeleteJob(ctx context.Context, req *miner_core.DeleteJobReq) (*miner_core.DeleteJobResp, error)
}

type JobServiceImpl struct {
	p JobServiceParam
}

type JobServiceParam struct {
	dig.In
	JobRepo  repo.JobRepo
	UserRepo repo.UserRepo
}

func NewJobService(p JobServiceParam) JobService {
	return &JobServiceImpl{
		p: p,
	}
}

func (s *JobServiceImpl) QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq) (r *miner_core.QueryJobListResp, err error) {
	jobs, err := s.p.JobRepo.QueryJobList(ctx, converter.BuildQueryJobListReq(req))
	if err != nil {
		logger.CtxErrorf(ctx, "s.p.JobRepo.QueryJobList failed, err = %v", err)
		return nil, err
	}

	userIDs := gslice.Union(
		gslice.Map(jobs, func(job domain.JobDO) int64 {
			return job.CreatedBy
		}),
		gslice.Map(jobs, func(job domain.JobDO) int64 {
			return job.UpdatedBy
		}),
	)

	userIDs2DO, err := s.p.UserRepo.QueryUser(ctx, domain.QueryUserReqDO{
		IDs: userIDs,
	})
	if err != nil {
		logger.CtxErrorf(ctx, "s.p.UserRepo.QueryUser failed, err = %v", err)
		return nil, err
	}
	dtos := converter.JobDOs2DTOs(jobs, userIDs2DO)

	return &miner_core.QueryJobListResp{
		JobList: dtos,
		Total:   int64(len(dtos)),
	}, nil
}

func (s *JobServiceImpl) CreateJob(ctx context.Context, req *miner_core.CreateJobReq) (r *miner_core.CreateJobResp, err error) {
	token := req.GetToken()
	claims, err := jwt.ValidateToken(ctx, token)
	if err != nil {
		logger.CtxErrorf(ctx, "jwt.ValidateToken failed, err = %v", err)
		return nil, err
	}
	userID, err := strconv.ParseInt(claims["user_id"], 10, 64)
	if err != nil {
		logger.CtxErrorf(ctx, "strconv.ParseInt failed, err = %v", err)
		return nil, err
	}
	id, err := s.p.JobRepo.CreateJob(ctx, &domain.JobDO{
		Name:        req.Name,
		Description: gptr.Indirect(req.Description),
		CreatedBy:   userID,
		UpdatedBy:   userID,
	})
	if err != nil {
		logger.CtxErrorf(ctx, "s.p.JobRepo.CreateJob failed, err = %v", err)
		return nil, err
	}
	return &miner_core.CreateJobResp{
		Id: id,
	}, nil
}

func (s *JobServiceImpl) DeleteJob(ctx context.Context, req *miner_core.DeleteJobReq) (*miner_core.DeleteJobResp, error) {
	token := req.GetToken()
	claims, err := jwt.ValidateToken(ctx, token)
	if err != nil {
		logger.CtxErrorf(ctx, "jwt.ValidateToken failed, err = %v", err)
		return nil, err
	}
	userID, err := strconv.ParseInt(claims["user_id"], 10, 64)
	if err != nil {
		logger.CtxErrorf(ctx, "strconv.ParseInt failed, err = %v", err)
		return nil, err
	}
	list, err := s.p.JobRepo.QueryJobList(ctx, &domain.QueryJobListReqDO{
		PageNum:  1,
		PageSize: 1,
		ID:       gptr.Of(req.Id),
	})
	if err != nil {
		logger.CtxErrorf(ctx, "s.p.JobRepo.QueryJobList failed, err = %v", err)
		return nil, err
	}
	if len(list) == 0 {
		logger.CtxErrorf(ctx, "job not found, id = %d", req.Id)
		return nil, biz_error.JobNotFoundError
	}
	if list[0].CreatedBy != userID {
		logger.CtxErrorf(ctx, "job that you want to delete is not yours, id = %d", req.Id)
		return nil, biz_error.DeleteJobError
	}
	err = s.p.JobRepo.DeleteJob(ctx, req.Id)
	if err != nil {
		logger.CtxErrorf(ctx, "s.p.JobRepo.DeleteJob failed, err = %v", err)
		return nil, err
	}
	return &miner_core.DeleteJobResp{}, nil
}
