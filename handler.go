package main

import (
	"context"
	miner_core "github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, req *miner_core.GetItemReq) (resp *miner_core.GetItemResp, err error) {
	// TODO: Your code here...
	return
}
