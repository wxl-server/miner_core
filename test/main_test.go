package test

import (
	"context"
	"miner_core/sal/config"
	"testing"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	Init(ctx)
}

func Init(ctx context.Context) {
	config.InitLocalConfig(ctx, "../conf/prod.yaml")
}
