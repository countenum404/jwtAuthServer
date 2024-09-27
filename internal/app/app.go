package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"jwtAuth/internal/api"
	"jwtAuth/internal/service/token"
	"jwtAuth/internal/service/user"
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
		token.Module,
		user.Module,
		postgres.Module,
		storage.Module,
		api.Module,
	).Run()
}
