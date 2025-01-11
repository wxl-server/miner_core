package app

import (
	"go.uber.org/dig"
	"miner_core/service"
)

type App struct {
	P Param
}

type Param struct {
	dig.In

	JobService service.JobService
}

func NewApp(p Param) *App {
	return &App{
		P: p,
	}
}
