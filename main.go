package main

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/qcq1/common/render"
	"go.uber.org/dig"
	"miner_core/app"
	"miner_core/sal/config"
	"miner_core/sal/dao"
	"miner_core/sal/id_generator"
	"miner_core/service"
)

var (
	initCtx   = context.Background()
	container = dig.New()
)

func main() {
	initContainer()

	mustInvoke(runServer)
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
		mustProvide(NewHandler)
	}
}

func mustProvide(constructor interface{}, opts ...dig.ProvideOption) {
	if err := container.Provide(constructor); err != nil {
		logger.Errorf("container provide failed, err = %v, constructor = %v", err, render.Render(constructor))
		panic(err)
	}
}

func mustInvoke(function interface{}, opts ...dig.InvokeOption) {
	if err := container.Invoke(function); err != nil {
		logger.Errorf("container invoke failed, err = %v, function = %v", err, render.Render(function))
		panic(err)
	}
}
