package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"jwtAuth/internal/api"
	"jwtAuth/internal/service"
	"jwtAuth/internal/storage"
	"jwtAuth/internal/storage/postgres"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	gin.SetMode(gin.DebugMode)
	fx.New(
		postgres.Module,
		storage.Module,
		service.Module,
		api.Module,
	).Run()
}
