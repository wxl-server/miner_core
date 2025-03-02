package service

import (
	"context"

	"github.com/wxl-server/idl_gen/kitex_gen/miner_core"
	"go.uber.org/dig"
)

type TaskService interface {
	RunTask(ctx context.Context, req *miner_core.RunTaskReq) (*miner_core.RunTaskResp, error)
}

type TaskServiceImpl struct {
	p TaskServiceParam
}

type TaskServiceParam struct {
	dig.In
}

func NewTaskService(p TaskServiceParam) TaskService {
	return &TaskServiceImpl{
		p: p,
	}
}

func (t TaskServiceImpl) RunTask(ctx context.Context, req *miner_core.RunTaskReq) (*miner_core.RunTaskResp, error) {
	//TODO implement me
	panic("implement me")
}
