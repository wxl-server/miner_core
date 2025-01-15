package main

import (
	"context"
	"miner_core/app"
	"miner_core/repo"
	"miner_core/sal/config"
	"miner_core/sal/dao"
	"miner_core/sal/id_generator"
	"miner_core/service"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/qcq1/common/render"
	"github.com/qcq1/common/wxl_cluster"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core/minercore"
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
		mustProvide(config.InitAppConfig)
	}

	// db
	{
		mustInvoke(dao.InitDB)
	}

	// id generator
	{
		mustProvide(id_generator.InitIDGenerator)
	}

	// dal
	{
		mustProvide(dao.NewJobDal)
	}

	// repo
	{
		mustProvide(repo.NewJobRepo)
	}

	// service
	{
		mustProvide(service.NewJobService)
	}

	// app
	{
		mustProvide(app.NewApp)
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
