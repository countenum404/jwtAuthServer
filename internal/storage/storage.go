package storage

import (
	"go.uber.org/fx"
	"jwtAuth/internal/storage/postgres"
)

var Module = fx.Module("storage",
	fx.Provide(fx.Annotate(postgres.NewUserStorage, fx.As(new(UserStorage)))),
	fx.Provide(fx.Annotate(postgres.NewTokenStorage, fx.As(new(TokenStorage)))),
)

type UserStorage interface {
	GetUser(username, password string) bool
}

type TokenStorage interface {
	GetRefreshToken(guid string) (string, error)
	SaveRefreshToken(guid, refreshToken, sessionId string) (string, error)
}
