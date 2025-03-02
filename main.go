package main

import (
	"context"
	"miner_core/repo"
	"miner_core/sal/config"
	"miner_core/sal/dao"
	"miner_core/service"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/common/reader"
	"github.com/wxl-server/common/render"
	"github.com/wxl-server/common/wxl_cluster"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_core/minercore"
	"go.uber.org/dig"
)

var (
	initCtx   = context.Background()
	container = dig.New()
)

func main() {
	initContainer()

	wxl_cluster.NewServer(minercore.NewServer, handler, "miner_core", 8889)
}

func initContainer() {
	// context
	{
		mustProvide(func() context.Context { return initCtx })
	}

	// config
	{
		mustProvide(reader.InitAppConfig[config.AppConfig])
	}

	// db
	{
		mustInvoke(dao.InitDB)
	}

	// repo
	{
		mustProvide(repo.NewJobRepo)
		mustProvide(repo.NewUserRepo)
	}

	// service
	{
		mustProvide(service.NewJobService)
		mustProvide(service.NewUserService)
		mustProvide(service.NewIndicatorService)
		mustProvide(service.NewTaskService)
	}

	// handler
	{
		mustInvoke(NewHandler)
	}
}

func mustProvide(constructor interface{}, opts ...dig.ProvideOption) {
	if err := container.Provide(constructor, opts...); err != nil {
		logger.Errorf("container provide failed, err = %v, constructor = %v", err, render.Render(constructor))
		panic(err)
	}
}

func mustInvoke(function interface{}, opts ...dig.InvokeOption) {
	if err := container.Invoke(function, opts...); err != nil {
		logger.Errorf("container invoke failed, err = %v, function = %v", err, render.Render(function))
		panic(err)
	}
}
