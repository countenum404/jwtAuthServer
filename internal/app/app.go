package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"jwtAuth/internal/api"
	tokenService "jwtAuth/internal/service/token"
	userService "jwtAuth/internal/service/user"
	"jwtAuth/internal/storage/postgres"
	tokenStorage "jwtAuth/internal/storage/token"
	userStorage "jwtAuth/internal/storage/user"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	gin.SetMode(gin.DebugMode)
	fx.New(
		postgres.Module,

		tokenService.Module,
		tokenStorage.Module,

		userService.Module,
		userStorage.Module,

		api.Module,
	).Run()
}
