package service

import (
	"context"
	"miner_core/common/consts"

	"github.com/wxl-server/common/json"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_core"
	"go.uber.org/dig"
)

type IndicatorService interface {
	QueryIndicatorList(ctx context.Context, req *miner_core.QueryIndicatorListReq) (*miner_core.QueryIndicatorListResp, error)
}

type IndicatorServiceImpl struct {
	p IndicatorServiceParam
}

type IndicatorServiceParam struct {
	dig.In
}

func NewIndicatorService(p IndicatorServiceParam) IndicatorService {
	return &IndicatorServiceImpl{
		p: p,
	}
}

func (i IndicatorServiceImpl) QueryIndicatorList(ctx context.Context, req *miner_core.QueryIndicatorListReq) (*miner_core.QueryIndicatorListResp, error) {
	resp := json.UnmarshalWithoutError[*miner_core.QueryIndicatorListResp](consts.Indicators)
	return resp, nil
}
