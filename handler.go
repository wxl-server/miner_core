package main

import (
	"context"
	"github.com/qcq1/common/json"

	"github.com/bytedance/gopkg/util/logger"
	miner_core "github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, req *miner_core.GetItemReq) (resp *miner_core.GetItemResp, err error) {
	logger.CtxInfof(ctx, "GetItem called, req = %v", json.MarshalWithoutError[string](req))
	return &miner_core.GetItemResp{
		Item: &miner_core.Item{
			Id:    1,
			Title: "test",
			Stock: 100,
		},
		BaseResp: nil,
	}, nil
}
