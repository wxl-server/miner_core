package main

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	miner_core "github.com/qcq1/rpc_miner_core/kitex_gen/miner_core/itemservice"
	"miner_core/sal/config"
)

func main() {
	ctx := context.Background()
	Init(ctx)

	svr := miner_core.NewServer(new(ItemServiceImpl))
	err := svr.Run()
	if err != nil {
		logger.CtxErrorf(ctx, "[Init] server run failed, err = %v", err)
		panic(err)
	}
}

func Init(ctx context.Context) {
	config.InitLocalConfig(ctx, "conf/prod.yaml")
}
